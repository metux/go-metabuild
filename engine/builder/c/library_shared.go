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

	cc := compiler.NewCCompiler(b.Compiler, b.TempDir())

	args := compiler.CompilerArg{
		Sources:    b.Parent.Sources(),
		PkgImports: b.Parent.AllImports(),
		Defines:    append(b.CDefs, b.CDefines()...),
		Output:     b.OutputFile(),
		Flags:      b.CFlags,
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
