package c

import (
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/compiler"
)

type BuilderCLibraryStatic struct {
	CommonCBuilder
}

func (b BuilderCLibraryStatic) JobRun() error {
	ci := b.BuildConf.CompilerInfo(b.ForBuild(), b.CompilerLang())
	cc := compiler.NewCCompiler(ci, b.TempDir())

	args := compiler.CompilerArg{
		Sources:    b.Sources(),
		PkgImports: b.Parent.AllImports(),
		Defines:    b.CDefines(),
		Output:     b.RequiredEntryStr(target.KeyFile),
		Flags:      b.CFlags(),
	}

	if err := cc.CompileLibraryStatic(args); err != nil {
		return err
	}

	b.InstallPkgFileAuto()

	return nil
}
