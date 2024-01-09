package base

import (
	"os/exec"

	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
)

func (b BaseBuilder) Exec(cmdline []string, wd string) error {
	if b.EntryBoolDef(target.KeyExecLog, true) {
		b.Logf("Exec: %s", cmdline)
	}
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Dir = wd
	cmd.Stdout = logWriter{Prefix: "out", Builder: &b}
	cmd.Stderr = logWriter{Prefix: "err", Builder: &b}
	return cmd.Run()
}

func (b BaseBuilder) ExecAbort(cmdline []string, wd string) {
	err := b.Exec(cmdline, wd)
	util.ErrPanicf(err, "exec failed: %s in %s", cmdline, wd)
}
