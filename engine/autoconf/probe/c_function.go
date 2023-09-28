package probe

import (
	"fmt"
)

type ProbeCFunction struct {
	ProbeCBase
}

func (p ProbeCFunction) Probe() error {
	f := p.Functions()
	src := p.CSource()
	for _, i := range f {
		src.Text.WriteString(fmt.Sprintf("extern int %s();\n", i))
	}
	for _, i := range f {
		src.MainBody.WriteString(fmt.Sprintf("    %s();\n", i))
	}
	return p.RunCheckCProg(src)
}

func MakeProbeCFunction(chk Check) ProbeInterface {
	p := ProbeCFunction{MakeProbeCBase(chk)}
	p.SetIdList(p.Functions())
	return p
}
