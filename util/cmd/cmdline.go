package cmd

import (
	"github.com/buildkite/shellwords"
	"os"
)

func StrCmdline(s string) []string {
	if s == "" {
		return []string{}
	}

	out, _ := shellwords.Split(s)
	return out
}

func EnvCmdline(envvar string) []string {
	return StrCmdline(os.Getenv(envvar))
}
