package c

import (
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/compiler"
)

type BuilderCLibraryStatic struct {
	CommonCBuilder
}

func (b BuilderCLibraryStatic) JobRun() error {
	cc := compiler.NewCCompiler(b.Compiler, b.TempDir())

	args := compiler.CompilerArg{
		Sources:    b.Parent.Sources(),
		PkgImports: b.Parent.AllImports(),
		Defines:    append(b.Parent.CDefines(), b.CDefines()...),
		Output:     b.RequiredEntryStr(target.KeyFile),
		Flags:      append(b.Parent.CFlags(), b.CFlags()...),
	}

	if err := cc.CompileLibraryStatic(args); err != nil {
		return err
	}

	b.InstallPkgFileAuto()

	return nil
}
