package c

import (
	"path/filepath"

	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/compiler"
)

type BuilderCLibraryShared struct {
	CommonCBuilder
}

func (b BuilderCLibraryShared) JobRun() error {
	soFile := b.OutputFile()

	ci := b.BuildConf.CompilerInfo(b.ForBuild(), b.CompilerLang())
	cc := compiler.NewCCompiler(ci, b.TempDir())

	args := compiler.CompilerArg{
		Sources:    b.Parent.Sources(),
		PkgImports: b.Parent.AllImports(),
		Defines:    append(b.Parent.CDefines(), b.CDefines()...),
		Output:     b.OutputFile(),
		Flags:      append(b.Parent.CFlags(), b.CFlags()...),
		DllName:    b.RequiredEntryStr(target.KeyName),
	}

	if err := cc.CompileLibraryShared(args); err != nil {
		return err
	}

	bname := filepath.Base(soFile)

	if b.InstallPkgFileAuto() {
		b.WritePkgMeta(bname+".sodep", cc.BinaryInfo(soFile).DependsInfo())
		b.WritePkgMeta(bname+".trigger", "activate-noawait ldconfig")
	}

	return nil
}
