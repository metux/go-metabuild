// This package contains the actual probes for auto configuration
package probe

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/check"
)

type Key = spec.Key
type Check = spec.Check

func mkProbe(chk Check) ProbeInterface {
	switch Key(chk.Type()) {
	case check.KeyCHeader:
		return MakeProbeCHeader(chk)
	case check.KeyCFunction:
		return MakeProbeCFunction(chk)
	case check.KeyCType:
		return MakeProbeCType(chk)
	case check.KeyPkgConfig:
		return MakeProbePkgConfig(chk)
	case check.KeyCCompiler:
		return MakeProbeCCompiler(chk)
	case check.KeyCXXCompiler:
		return MakeProbeCXXCompiler(chk)
	case check.KeyTargetDistro:
		return MakeProbeTargetDistro(chk)
	case check.KeyGitDescribe:
		return MakeGitDescribe(chk)
	case check.KeyI18nLinguas:
		return MakeI18nLinguas(chk)
	}
	return nil
}

// FIXME: store error in cache ?
func Probe(chk Check) bool {
	probe := mkProbe(chk)
	if probe == nil {
		panic("unsupported check: " + chk.Type())
	}

	chk.Logf("checking: %s", chk)

	// check for cached value ?
	if cached, cacheval := chk.Cached(); cached {
		chk.Logf("cached: %t", cacheval)
		return cacheval
	}

	if err := probe.Probe(); err != nil {
		chk.Logf("check failed: %s\n", err)
		chk.Done(false)
		return false
	}

	chk.Done(true)
	return true
}
