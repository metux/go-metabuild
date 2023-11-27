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

func (b BuilderDataMisc) JobPrepare(id string) error {
	return nil
}

func MakeDataMisc(o spec.TargetObject, id string) BuilderDataMisc {
	return BuilderDataMisc{base.BaseBuilder{o, id}}
}
