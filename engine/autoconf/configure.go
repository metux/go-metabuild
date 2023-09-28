package autoconf

import (
	"fmt"

	"github.com/metux/go-metabuild/engine/autoconf/probe"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/spec/check"
	"github.com/metux/go-metabuild/spec/global"
)

func RunConfigure(cf global.Global) error {
	for _, chk := range cf.GetChecks() {
		if err := runCheck(cf, chk); err != nil {
			return err
		}
	}
	return RunConfigureFeatures(cf)
}

func runCheck(cf global.Global, chk spec.Check) error {
	res := probe.Probe(chk)
	if chk.IsMandatory() && !res {
		return fmt.Errorf("%w: %s", ErrCheckFailedMandatory, chk)
	}

	// store explicit C defines (fixme: obsoleted ?)
	if res {
		for _, sym := range chk.GetDefines() {
			cf.EntryPutStr(global.KeyCheckedCDefines.AppendStr(sym), "1")
		}
	}

	// set yes/no specific symbols
	flags := chk.BuildConf.Flags(chk.ForBuild())
	flags.EntryStrListAppendList(check.KeyCDefines, chk.YesNoStrList(res, check.KeyCDefines))
	flags.EntryStrListAppendList(check.KeyCCFlags, chk.YesNoStrList(res, check.KeyCCFlags))
	flags.EntryStrListAppendList(check.KeyCLDFlags, chk.YesNoStrList(res, check.KeyCLDFlags))

	// store .config flags
	if cf := chk.EntryStr(buildconf.KeyCheckConfig); cf != "" {
		chk.BuildConf.ConfigBool(cf, res)
	}

	return nil
}