package gen

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
)

type GlibMarshal struct {
	base.BaseBuilder
}

func (b GlibMarshal) JobRun() error {

	c1 := append(b.BuilderCmd(),
		"--prefix",
		b.RequiredEntryStr(target.KeyResourceName),
		b.RequiredSourceAbs(),
		"--output")

	b.ExecAbort(append(c1, b.RequiredEntryStr(target.KeyOutputCHeader), "--header"), "")
	b.ExecAbort(append(c1, b.RequiredEntryStr(target.KeyOutputCSource), "--body"), "")

	return nil
}

func MakeGlibMarshal(o spec.TargetObject, id string) GlibMarshal {
	return GlibMarshal{base.BaseBuilder{o, id}}
}
