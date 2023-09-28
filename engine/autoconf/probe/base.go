package probe

type ProbeInterface interface {
	Init() error
	Probe() error
}

type ProbeBase struct {
	Check
}

func (cb ProbeBase) Init() error {
	return nil
}

func (cb ProbeBase) Logf(format string, args ...any) {
	cb.Check.Logf(format, args...)
}

func (cb ProbeBase) Id() string {
	return cb.Check.Id()
}

func MakeProbeBase(chk Check) ProbeBase {
	return ProbeBase{chk}
}
