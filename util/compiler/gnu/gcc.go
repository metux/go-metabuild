package gnu

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/compiler/base"
)

// FIXME: support multiple compiler types
// FIXME: move over to builders ?

type CCompilerGCC struct {
	CompilerInfo base.CompilerInfo
	cmdline      []string
	TempDir      string
}

// --- command line construction ---
// NOTE: these *need* pointer receiver

func (cc *CCompilerGCC) cmdCC(out string, src []string) {
	cc.cmdline = cc.CompilerInfo.Command
	cc.cmdline = append(cc.cmdline, "-o", out)
	cc.cmdline = append(cc.cmdline, src...)
}

func (cc *CCompilerGCC) cmdPkgImport(pkgimport []base.PkgConfigInfo) {
	// FIXME: filter out duplicate flags
	for _, pi := range pkgimport {
		if pi.WantStatic && pi.WantShared {
			panic("config error: cant have WantStatic && WantShared together")
		}

		if pi.WantStatic {
			cc.cmdline = append(cc.cmdline, pi.StaticCflags...)
			cc.cmdline = append(cc.cmdline, pi.StaticLdflags...)
		} else {
			cc.cmdline = append(cc.cmdline, pi.SharedCflags...)
			cc.cmdline = append(cc.cmdline, pi.SharedLdflags...)
		}
	}
}

func (cc *CCompilerGCC) cmdDefines(cdefs []string) {
	for _, x := range cdefs {
		cc.cmdline = append(cc.cmdline, "-D"+x)
	}
}

func (cc *CCompilerGCC) cmdShared() {
	cc.cmdline = append(cc.cmdline, "-fPIC", "-shared")
}

func (cc *CCompilerGCC) cmdSoName(soname string) {
	cc.cmdline = append(cc.cmdline, "-Wl,-soname,"+soname)
}

func (cc *CCompilerGCC) cmdAr(out string, objs []string) {
	cc.cmdline = append(append(cc.CompilerInfo.Archiver, "-rc", out), objs...)
}

func (cc *CCompilerGCC) cmdAppend(args ...string) {
	cc.cmdline = append(cc.cmdline, args...)
}

func (cc CCompilerGCC) cmdExec() error {
	log.Println("CC cmd", cc.cmdline)
	out, err := cmd.RunOut(cc.cmdline, true)
	if out != "" {
		log.Println(out)
	}
	return err
}

func (cc CCompilerGCC) CompileExecutable(args base.CompilerArg) error {
	if args.Output == "" {
		return fmt.Errorf("CompileExe: no output file")
	}

	cc.cmdCC(args.Output, args.Sources)
	cc.cmdPkgImport(args.PkgImports)
	cc.cmdDefines(args.Defines)
	return cc.cmdExec()
}

func (cc CCompilerGCC) CompileLibraryStatic(args base.CompilerArg) error {
	objs := []string{}
	cmdlines := [][]string{}

	// FIXME: move this into cc.cmdCompileObject()
	for _, cfile := range args.Sources {
		ofile := cc.TempDir + "/" + cfile + ".o"
		os.MkdirAll(filepath.Dir(ofile), 0755)
		objs = append(objs, ofile)
		cc2 := CCompilerGCC{CompilerInfo: cc.CompilerInfo, TempDir: cc.TempDir}
		cc2.cmdline = cc.CompilerInfo.Command
		cc2.cmdDefines(args.Defines)
		cc2.cmdPkgImport(args.PkgImports)
		cc2.cmdAppend("-o", ofile, "-c", cfile)
		cmdlines = append(cmdlines, cc2.cmdline)
	}

	outs, errs := cmd.RunGroup(cmdlines)
	for idx, _ := range outs {
		if outs[idx] != "" {
			log.Println(outs[idx])
		}
		if errs[idx] != nil {
			log.Printf("Compile object %s failed: %s", objs[idx], errs[idx])
			return errs[idx]
		}
	}

	cc.cmdAr(args.Output, objs)
	return cc.cmdExec()
}

func (cc CCompilerGCC) CompileLibraryShared(args base.CompilerArg) error {
	cc.cmdCC(args.Output, args.Sources)
	cc.cmdPkgImport(args.PkgImports)
	cc.cmdDefines(args.Defines)
	cc.cmdShared()
	cc.cmdSoName(args.DllName)
	return cc.cmdExec()
}

// FIXME: need to probe objdump
func (cc CCompilerGCC) elfDepends(src string) []string {
	libs, err := ELFDepends([]string{}, src)
	if err != nil {
		log.Println("CCompiler: ELFDepends failed", err)
	}
	return libs
}

func (cc CCompilerGCC) toolObjdump() []string {
	return []string{}
}

func (cc CCompilerGCC) binaryArch(fn string) string {
	arch, err := ELFArch(cc.toolObjdump(), fn)
	if err != nil {
		log.Println("CCompiler: ELFDepends failed", err)
	}
	return arch
}

// FIXME: check type and format
func (cc CCompilerGCC) BinaryInfo(fn string) base.BinaryFileInfo {
	return base.BinaryFileInfo{
		Filename:   fn,
		Depends:    cc.elfDepends(fn),
		BinaryArch: cc.binaryArch(fn),
	}
}

func NewCCompiler(ci base.CompilerInfo, tempdir string) base.CCompiler {
	return CCompilerGCC{
		CompilerInfo: ci,
		TempDir:      tempdir,
	}
}
