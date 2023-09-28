package probe

import (
	"fmt"
)

type ProbeCType struct {
	ProbeCBase
}

func (p ProbeCType) Probe() error {
	src := p.CSource()
	for x, i := range p.Types() {
		src.Text.WriteString(fmt.Sprintf("%s test_%d;\n", i, x))
	}
	return p.RunCheckCProg(src)
}

func MakeProbeCType(chk Check) ProbeInterface {
	p := ProbeCType{MakeProbeCBase(chk)}
	p.SetIdList(p.Types())
	return p
}
