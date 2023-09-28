package probe

import (
	"fmt"
	"strings"

	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/distro"
)

type ProbeTargetDistro struct {
	ProbeBase
}

// FIXME: support crosscompile
// FIXME: support sysroot
func (p ProbeTargetDistro) probeDEB() error {
	out, err := cmd.RunOutOne([]string{"dpkg", "--print-architecture"}, true)
	if err != nil {
		return err
	}
	p.BuildConf.SetTargetDistArch(out)
	return nil
}

// FIXME: support RPM
func (p ProbeTargetDistro) probeRPM() error {
	p.Logf("WARN: RPM not implemented yet")
	return nil
}

// FIXME: support other OS'es besides Linux
// FIXME: check globalconf for supported distros ? (id vs idlike)
// FIXME: allow overriding via env
func (p ProbeTargetDistro) Probe() error {
	inf, err := distro.DistroDetect("")
	if err != nil {
		return fmt.Errorf("failed to detect distro -- missing /etc/os-release")
	}

	p.Check.BuildConf.SetTargetDist(strings.ToLower(inf.IdLike))

	targetdist := p.Check.BuildConf.TargetDistro()
	pkgfmt := targetdist.PackageFormat()

	switch pkgfmt {
	case distro.PkgFormatDEB:
		return p.probeDEB()
	case distro.PkgFormatRPM:
		return p.probeRPM()
	}

	p.Logf("WARN: cant detect packaging format -- %s", pkgfmt)
	return nil
}

func MakeProbeTargetDistro(chk Check) ProbeTargetDistro {
	return ProbeTargetDistro{MakeProbeBase(chk)}
}
