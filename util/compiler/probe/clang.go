package probe

import (
	cmdrun "github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/compiler/base"
)

func probeClang(ci *base.CompilerInfo) bool {
	ci.Id = base.CompilerGCC
	ci.Machine = base.ParseMachine(ci.RunCatchOut("-dumpmachine"))
	ci.Version = ci.RunCatchOut("-dumpversion")
	ci.Archiver = cmdrun.StrCmdline(probeArCmd(ci))
	return ci.Version != ""
}
