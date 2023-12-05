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
	if err := RunConfigureFeatures(cf); err != nil {
		return err
	}

	// run post-config stage
	cf.PostConfig()

	return nil
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

	k := chk.YesNoKey(res, check.KeyPkgConfAdd)
	b := chk.ForBuild()
	for _, i := range chk.EntryKeys(k) {
		id := string(i)
		chk.BuildConf.SetPkgConfig(b, id, buildconf.PkgconfLoad(chk.EntrySpec(k.Append(i)), id))
	}

	// store .config flags
	if cf := chk.EntryStr(buildconf.KeyCheckConfig); cf != "" {
		chk.BuildConf.ConfigBool(cf, res)
	}

	return nil
}
