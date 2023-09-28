package doc

// FIXME: automatic compression

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
)

type BuilderDocMisc struct {
	base.BaseBuilder
}

func (b BuilderDocMisc) JobRun() error {
	if !b.WantInstall() {
		return nil
	}

	installdir := b.InstallDir()
	fmode := b.InstallPerm()
	compress := b.EntryStr("compress")
	for _, src := range b.Sources() {
		b.InstallPkgFileCompressed(src, installdir, fmode, compress)
	}
	return nil
}

// FIXME: move this to base builder ?
func (b BuilderDocMisc) JobPrepare(id string) error {
	b.LoadTargetDefaults(spec.Key(b.Type()))
	return nil
}

func MakeDocMisc(o spec.TargetObject, id string) BuilderDocMisc {
	return BuilderDocMisc{base.BaseBuilder{o, id}}
}
