package gen

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/cmd"
)

type GlibResource struct {
	base.BaseBuilder
}

func (b GlibResource) JobRun() error {
	src := b.Sources()[0]
	cheader := b.RequiredEntryStr(target.KeyOutputCHeader)
	csource := b.RequiredEntryStr(target.KeyOutputCSource)
	gresource := b.RequiredEntryStr(target.KeyOutputGResource)
	srcdir := b.EntryStr(target.KeyResourceDir)

	c1 := []string{"glib-compile-resources", src}
	if srcdir != "" {
		c1 = append(c1, "--sourcedir="+srcdir)
	}

	_, err := cmd.RunOutOne(append(c1, "--target="+cheader, "--generate-header"), true)
	util.ErrPanicf(err, "failed generating header")

	_, err = cmd.RunOutOne(append(c1, "--target="+csource, "--generate-source"), true)
	util.ErrPanicf(err, "failed generating source")

	_, err = cmd.RunOutOne(append(c1, "--target="+gresource), true)
	util.ErrPanicf(err, "failed generating gresource")

	return nil
}

func MakeGlibResource(o spec.TargetObject, id string) GlibResource {
	return GlibResource{base.BaseBuilder{o, id}}
}
