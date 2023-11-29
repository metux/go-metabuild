package probe

import (
	"path/filepath"

	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/spec/check"
)

type I18nLinguas struct {
	ProbeBase
}

// FIXME: shortcut notation ?
func (p I18nLinguas) Probe() error {
	podir := p.EntryStr(check.KeyI18nPoDir)
	if podir == "" {
		podir = "po"
	}

	linguas, err := filepath.Glob(podir + "/*.po")
	if err != nil {
		return err
	}

	for idx, f := range linguas {
		f := filepath.Base(f)
		ext := filepath.Ext(f)
		if ext != "" {
			f = f[:len(f)-len(ext)]
		}
		linguas[idx] = f
	}

	p.Logf("detected linguas: %s", linguas)
	p.BuildConf.EntryPutStrList(buildconf.KeyLinguas, linguas)

	return nil
}

func MakeI18nLinguas(chk Check) ProbeInterface {
	return I18nLinguas{MakeProbeBase(chk)}
}
