package probe

import (
	"github.com/metux/go-metabuild/util/compiler"
)

type ProbeCXXCompiler struct {
	ProbeBase
}

func (p ProbeCXXCompiler) Probe() error {

	infoTarget, infoHost, err := compiler.DetectCXX()

	if err == nil {
		// store target compiler settings
		p.Check.BuildConf.SetCompilerInfo(false, infoTarget)
		p.Check.BuildConf.SetCompilerInfo(true, infoHost)
	}

	return err
}

func MakeProbeCXXCompiler(chk Check) ProbeInterface {
	return ProbeCXXCompiler{MakeProbeBase(chk)}
}
