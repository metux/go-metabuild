package c

import (
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
		DllName:    b.RequiredEntryStr(target.KeyName),
	}

	if err := cc.CompileLibraryShared(args); err != nil {
		return err
	}

	if b.InstallPkgFileAuto() {
		b.WritePkgMeta(soFile+".sodep", cc.BinaryInfo(soFile).DependsInfo())
		b.WritePkgMeta(soFile+".trigger", "activate-noawait ldconfig")
	}

	return nil
}
