package base

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/metux/go-metabuild/util"
)

func (b BaseBuilder) Exec(cmdline []string, wd string) (string, error) {
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Dir = wd
	out, err := cmd.CombinedOutput()
	if err != nil {
		b.Logf("Command error for: %s\n", cmdline)
		b.Logf(">> %s\n", out)
	}
	return strings.TrimSpace(fmt.Sprintf("%s", out)), err
}

func (b BaseBuilder) ExecAbort(cmdline []string, wd string) string {
	out, err := b.Exec(cmdline, wd)
	util.ErrPanicf(err, "exec failed: %s in %s", cmdline, wd)
	return out
}
