package probe

import (
	"github.com/metux/go-metabuild/util/compiler"
)

type ProbeCCompiler struct {
	ProbeBase
}

func (p ProbeCCompiler) Probe() error {

	infoTarget, infoHost, err := compiler.DetectCC()

	if err == nil {
		// store target compiler settings
		p.Check.BuildConf.SetCompilerInfo(false, infoTarget)
		p.Check.BuildConf.SetCompilerInfo(true, infoHost)
	}

	return err
}

func MakeProbeCCompiler(chk Check) ProbeInterface {
	return ProbeCCompiler{MakeProbeBase(chk)}
}
