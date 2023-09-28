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
		Sources: b.Parent.Sources(),
		PkgImports: b.Parent.AllImports(),
		Defines: append(b.CDefs, b.CDefines()...),
		Output:  b.RequiredEntryStr(target.KeyFile),
	}

	if err := cc.CompileLibraryStatic(args); err != nil {
		return err
	}

	b.InstallPkgFileAuto()

	return nil
}
