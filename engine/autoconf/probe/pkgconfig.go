package probe

import (
	"github.com/metux/go-metabuild/spec/check"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/compiler"
)

type ProbePkgConfig struct {
	ProbeBase
}

func (p ProbePkgConfig) Probe() error {
	m := p.EntryStrMap(check.KeyPkgConfig)

	forBuild := p.Check.ForBuild()
	envvar := util.ValIf(forBuild, "HOST_PKG_CONFIG", "PKG_CONFIG")
	cmdline := cmd.EnvCmdline(envvar)
	if len(cmdline) == 0 {
		p.Logf("$%s not defined. assuming pkg-config", envvar)
		cmdline = cmd.StrCmdline("pkg-config")
	}

	var res error

	for id, query := range m {
		if info, err := compiler.PkgConfigQuery(query, cmdline); err == nil {
			p.Logf("pkgconf found: %s => %s", id, query)
			p.Check.BuildConf.SetPkgConfig(forBuild, string(id), info)
		} else {
			p.Logf("pkgconf missing: %s => %s (%s)", id, query, err)
			res = err
		}
	}

	return res
}

func MakeProbePkgConfig(chk Check) ProbePkgConfig {
	return ProbePkgConfig{MakeProbeBase(chk)}
}
