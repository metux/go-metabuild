package c

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util/compiler"
)

type BuilderCExecutable struct {
	BaseCBuilder
}

func (b BuilderCExecutable) JobRun() error {
	ci := b.BuildConf.CompilerInfo(b.ForBuild(), compiler.LangC)

	fn := b.OutputFile()

	cc := compiler.NewCCompiler(ci, b.TempDir())
	args := compiler.CompilerArg{
		Sources:    b.Sources(),
		PkgImports: b.AllImports(),
		Defines:    b.CDefines(),
		Output:     fn,
	}

	if err := cc.CompileExecutable(args); err != nil {
		return err
	}

	if b.InstallPkgFileAuto() {
		// FIXME: need to blacklist our own libs
		b.WritePkgMeta(fn+".sodep", cc.BinaryInfo(fn).DependsInfo())
	}

	return nil
}

func (b BuilderCExecutable) JobPrepare(id string) error {
	b.LoadTargetDefaults(spec.Key(b.Type()))
	return nil
}

func MakeBuilderCExecutable(o spec.TargetObject, id string) BuilderCExecutable {
	return BuilderCExecutable{NewBaseCBuilder(o, id)}
}
