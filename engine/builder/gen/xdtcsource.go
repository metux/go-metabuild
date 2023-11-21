package gen

import (
	"io/ioutil"

	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/fileutil"
	"github.com/metux/go-metabuild/util/strs"
)

type XdtCSource struct {
	base.BaseBuilder
}

func (b XdtCSource) JobRun() error {
	src := b.Sources()[0]
	cheader := b.RequiredEntryStr(target.KeyOutputCHeader)
	resname := b.RequiredEntryStr(target.KeyResourceName)

	raw, err := ioutil.ReadFile(src)
	util.ErrPanicf(err, "failed reading input: "+src)

	return fileutil.WriteText(cheader, strs.XdtCSource(string(raw), src, resname))
}

func MakeXdtCSource(o spec.TargetObject, id string) XdtCSource {
	return XdtCSource{base.BaseBuilder{o, id}}
}
