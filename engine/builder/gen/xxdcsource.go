package gen

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/strs"
)

type XxdCSource struct {
	base.BaseBuilder
}

func (b XxdCSource) JobRun() error {
	src := b.RequiredSourceAbs()
	cheader := b.RequiredEntryStr(target.KeyOutputCHeader)
	return strs.XXD(src, cheader)
}

func MakeXxdCSource(o target.TargetObject, id string) XxdCSource {
	return XxdCSource{base.BaseBuilder{o, id}}
}
