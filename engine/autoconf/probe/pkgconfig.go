package probe

import (
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/compiler"
)

type ProbePkgConfig struct {
	ProbeBase
}

func (p ProbePkgConfig) Probe() error {
	id := p.Id()
	if id == "" {
		return ErrNeedId
	}

	forBuild := p.Check.ForBuild()

	if pi := p.Check.BuildConf.PkgConfig(forBuild, id); pi.Name != "" {
		return nil
	}

	envvar := util.ValIf(forBuild, "HOST_PKG_CONFIG", "PKG_CONFIG")
	cmdline := cmd.EnvCmdline(envvar)
	if len(cmdline) == 0 {
		p.Logf("$%s not defined. assuming pkg-config", envvar)
		cmdline = cmd.StrCmdline("pkg-config")
	}

	// FIXME: should fetch static vs dyn into separate variables
	info, err := compiler.PkgConfigQuery(p.Check.EntryStr("pkg-config"), cmdline)
	p.Check.BuildConf.SetPkgConfig(forBuild, id, info)

	return err
}

func MakeProbePkgConfig(chk Check) ProbePkgConfig {
	return ProbePkgConfig{MakeProbeBase(chk)}
}
