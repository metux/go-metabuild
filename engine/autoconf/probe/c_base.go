package probe

import (
	"fmt"
	"os"

	"github.com/metux/go-metabuild/spec/check"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/compiler"
)

type ProbeCBase struct {
	ProbeBase
}

func (p ProbeCBase) RunCheckCProg(cs util.CSource) error {
	// check if host/target

	src := p.Check.TestProgFile("c")
	out := p.Check.TestProgFile("exe")

	defer os.Remove(src)
	defer os.Remove(out)

	if err := cs.Write(src); err != nil {
		return err
	}

	cc := compiler.NewCCompiler(p.Check.BuildConf.CompilerInfo(p.Check.ForBuild(), compiler.LangC), p.Check.TempDir())

	args := compiler.CompilerArg{
		Sources:    []string{src},
		PkgImports: []compiler.PkgConfigInfo{},
		Defines:    p.Check.EntryStrList("c/defines"),
		Output:     out,
	}

	if err := cc.CompileExecutable(args); err != nil {
		return fmt.Errorf("%w: %s", ErrCompileFailed, err)
	}

	return nil
}

func (p ProbeCBase) Headers() []string {
	return p.Check.EntryStrList(check.KeyCHeader)
}

func (p ProbeCBase) Functions() []string {
	return p.Check.EntryStrList(check.KeyCFunction)
}

func (p ProbeCBase) CSource() util.CSource {
	return util.CSource{Includes: p.Headers()}
}

func (p ProbeCBase) Types() []string {
	return p.Check.EntryStrList(check.KeyCType)
}

func MakeProbeCBase(chk Check) ProbeCBase {
	return ProbeCBase{MakeProbeBase(chk)}
}
