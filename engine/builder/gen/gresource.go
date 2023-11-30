package gen

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
)

type GlibResource struct {
	base.BaseBuilder
}

func (b GlibResource) JobRun() error {
	cheader := b.RequiredEntryStr(target.KeyOutputCHeader)
	csource := b.RequiredEntryStr(target.KeyOutputCSource)
	gresource := b.RequiredEntryStr(target.KeyOutputGResource)

	c1 := append(b.BuilderCmd(), b.RequiredSourceAbs())

	if srcdir := b.EntryStr(target.KeyResourceDir); srcdir != "" {
		c1 = append(c1, "--sourcedir="+srcdir)
	}

	b.ExecAbort(append(c1, "--target="+cheader, "--generate-header"), "")
	b.ExecAbort(append(c1, "--target="+csource, "--generate-source"), "")
	b.ExecAbort(append(c1, "--target="+gresource), "")

	return nil
}

func MakeGlibResource(o spec.TargetObject, id string) GlibResource {
	return GlibResource{base.BaseBuilder{o, id}}
}
