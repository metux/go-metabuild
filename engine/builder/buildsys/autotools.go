package buildsys

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/options"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/compiler"
	"github.com/metux/go-metabuild/util/fileutil"
)

type Autotools struct {
	base.BaseBuilder
}

func (b Autotools) stage(stage string, kEnable Key, kCommand Key, kWorkdir Key, kLog Key, args []string) error {
	b.Logf("running stage: %s", stage)
	if b.EntryBoolDef(kEnable, true) {
		return b.Exec(append(b.RequiredEntryStrList(kCommand), args...), b.EntryStr(kWorkdir))
	}
	return nil
}

func addPar(list []string, name string, val []string) []string {
	if len(val) == 0 {
		return list
	}
	return append(list, name+"="+strings.Join(val, " ")+"")
}

func (b Autotools) JobRun() error {

	destdir := filepath.Clean(b.EntryStr("install/destdir"))

	infCHost := b.BuildConf.CompilerInfo(true, compiler.LangC)
	infCBuild := b.BuildConf.CompilerInfo(false, compiler.LangC)
	infCxxHost := b.BuildConf.CompilerInfo(true, compiler.LangCxx)
	infCxxBuild := b.BuildConf.CompilerInfo(false, compiler.LangCxx)

	args := []string{
		"--host=" + infCHost.Machine.String(),
		"--build=" + infCBuild.Machine.String(),
	}

	args = append(args, b.RequiredEntryStrList(options.KeyAutoconfDirOpts)...)
	args = append(args, b.BuildFlags().EntryStrList(options.KeyAutoconfDirOpts)...)

	args = addPar(args, "CC", infCHost.Command)
	args = addPar(args, "AR", infCHost.Archiver)
	args = addPar(args, "LD", infCHost.Linker)
	args = addPar(args, "CXX", infCxxHost.Command)
	args = addPar(args, "HOST_CC", infCBuild.Command)
	args = addPar(args, "HOST_AR", infCBuild.Archiver)
	args = addPar(args, "HOST_LD", infCBuild.Linker)
	args = addPar(args, "HOST_CXX", infCxxBuild.Command)

	if err := b.stage("autogen",
		target.KeyAutotoolsAutogen,
		target.KeyAutotoolsAutogenCommand,
		target.KeyAutotoolsAutogenWorkDir,
		target.KeyAutotoolsAutogenLog,
		[]string{}); err != nil {
		return err
	}

	if err := b.stage("configure",
		target.KeyAutotoolsConfigure,
		target.KeyAutotoolsConfigureCommand,
		target.KeyAutotoolsConfigureWorkDir,
		target.KeyAutotoolsConfigureLog,
		args); err != nil {
		return err
	}

	if err := b.stage("clean",
		target.KeyAutotoolsClean,
		target.KeyAutotoolsCleanCommand,
		target.KeyAutotoolsCleanWorkDir,
		target.KeyAutotoolsCleanLog,
		[]string{}); err != nil {
		return err
	}

	if err := b.stage("build",
		target.KeyAutotoolsBuild,
		target.KeyAutotoolsBuildCommand,
		target.KeyAutotoolsBuildWorkDir,
		target.KeyAutotoolsBuildLog,
		[]string{}); err != nil {
		return err
	}

	if err := b.stage("install",
		target.KeyAutotoolsInstall,
		target.KeyAutotoolsInstallCommand,
		target.KeyAutotoolsInstallWorkDir,
		target.KeyAutotoolsInstallLog,
		[]string{}); err != nil {
		return err
	}

	// install files
	destFS := os.DirFS(destdir)
	installed := make(map[string]string)

	// FIXME: check if glob finds nothing
	globs := b.EntrySpec(target.KeyInstallGlob)
	for _, pname := range globs.Keys() {
		pkgpath := b.PkgPath(string(pname), "/")
		for _, ent := range globs.EntryStrList(pname) {
			found, _ := fs.Glob(destFS, fileutil.StripAbs(ent))
			for _, fn := range found {
				if _, ok := installed[fn]; ok {
					return fmt.Errorf("already installed in <%s>: %s", pname, fn)
				}
				idir := pkgpath + "/" + filepath.Dir(fn)
				os.MkdirAll(idir, 0755)
				fileutil.Copy(destdir+"/"+fn, idir)
				installed[fn] = string(pname)
			}
		}
	}

	// mark ignored
	for _, ent := range b.EntryStrList("install/ignore") {
		log.Println("ignore glob", ent)
		found, _ := fs.Glob(destFS, fileutil.StripAbs(ent))
		for _, fn := range found {
			if pname, ok := installed[fn]; ok {
				return fmt.Errorf("already installed in <%s>: %s", pname, fn)
			}
			installed[fn] = ""
		}
	}

	l := len(destdir) + 1
	err := filepath.Walk(destdir,
		func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if path != destdir && !info.IsDir() {
					path = path[l:]
					if _, ok := installed[path]; !ok {
						log.Println("missing: ", path)
					} else {
						log.Println("got: ", path)
					}
				}
			}
			return err
		})

	return err
}

func MakeAutotools(o spec.TargetObject, id string) Autotools {
	return Autotools{base.BaseBuilder{o, id}}
}
