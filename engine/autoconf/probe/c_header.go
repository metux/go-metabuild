package probe

type ProbeCHeader struct {
	ProbeCBase
}

func (p ProbeCHeader) Probe() error {
	src := p.CSource()
	return p.RunCheckCProg(src)
}

func MakeProbeCHeader(chk Check) ProbeInterface {
	p := ProbeCHeader{MakeProbeCBase(chk)}
	p.SetIdList(p.Headers())
	return p
}
