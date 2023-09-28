package data

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
)

type BuilderDataMisc struct {
	base.BaseBuilder
}

func (b BuilderDataMisc) JobRun() error {
	installdir := b.InstallDir()
	fmode := b.InstallPerm()
	for _, src := range b.Sources() {
		b.InstallPkgFile(src, installdir, fmode)
	}
	return nil
}

// FIXME: move this to base builder ?
func (b BuilderDataMisc) JobPrepare(id string) error {
	b.LoadTargetDefaults(spec.Key(b.Type()))
	return nil
}

func MakeDataMisc(o spec.TargetObject, id string) BuilderDataMisc {
	return BuilderDataMisc{base.BaseBuilder{o, id}}
}