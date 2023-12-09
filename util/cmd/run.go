package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunOut(cmdline []string, logErr bool) (string, error) {
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	if out, err := cmd.CombinedOutput(); err != nil {
		if logErr {
			log.Printf("Command error for: %s\n", cmdline)
			log.Printf(">> %s\n", out)
		}
		return fmt.Sprintf("%s", out), err
	} else {
		return fmt.Sprintf("%s", out), nil
	}
}

func RunGroup(cmdlines [][]string) ([]string, []error) {
	outs := make([]string, len(cmdlines))
	errs := make([]error, len(cmdlines))
	cmds := make([](*exec.Cmd), len(cmdlines))

	for x, y := range cmdlines {
		cmds[x] = exec.Command(y[0], y[1:]...)
		cmds[x].Stdout = os.Stdout
		cmds[x].Stderr = os.Stderr
		cmds[x].Start()
	}

	for x, _ := range cmdlines {
		errs[x] = cmds[x].Wait()
		// FIXME: capture output and stderr
		outs[x] = cmds[x].String()
	}

	return outs, errs
}

func RunOutOne(cmdline []string, logErr bool) (string, error) {
	out, err := RunOut(cmdline, logErr)
	return strings.TrimSpace(out), err
}

func RunOutCmd(cmdline []string, logErr bool) ([]string, error) {
	out, err := RunOut(cmdline, logErr)
	return StrCmdline(out), err
}
