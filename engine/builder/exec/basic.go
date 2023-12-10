package exec

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
)

type ExecBasic struct {
	base.BaseBuilder
}

func (b ExecBasic) JobRun() error {
	out, err := b.Exec(b.RequiredEntryStrList(target.KeyExecCommand),
		b.EntryStr(target.KeyExecWorkDir))

	if b.EntryBoolDef(target.KeyExecLog, true) {
		b.Logf("%s", out)
	}

	return err
}

func MakeExecBasic(o spec.TargetObject, id string) ExecBasic {
	return ExecBasic{base.BaseBuilder{o, id}}
}
