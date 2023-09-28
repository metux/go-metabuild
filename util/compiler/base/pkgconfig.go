package base

import (
	"strings"

	"github.com/buildkite/shellwords"

	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/fileutil"
)

// FIXME: how differenciate between static and dynamic ?
// FIXME: should we split it into two separate instances ?
type PkgConfigInfo struct {
	Id          string
	PkgSpec     string
	Name        string
	Version     string
	Description string

	Requires []string

	// FIXME: only-L only-I etc
	SharedCflags  []string
	SharedLdflags []string

	StaticCflags  []string
	StaticLdflags []string

	Variables map[string]string

	// These aren't read from .pc files, but may be set by other places
	Private    bool
	WantStatic bool
	WantShared bool
}

func (p PkgConfigInfo) Write(fn string) error {
	wr, err := fileutil.CreateWriter(fn)
	defer wr.Close()
	if err != nil {
		return err
	}

	wr.WriteField("Name", p.Name)
	wr.WriteField("Description", p.Description)
	wr.WriteField("Version", p.Version)
	wr.WriteField("Libs", strings.Join(p.SharedLdflags, " "))
	wr.WriteField("Cflags", strings.Join(p.SharedCflags, " "))
	for _, x := range p.Requires {
		wr.WriteField("Requires", x)
	}

	wr.WriteString("\n")

	for k, v := range p.Variables {
		wr.WriteVar(k, v)
	}

	return nil
}

func (p PkgConfigInfo) Valid() bool {
	return p.PkgSpec != ""
}

// FIXME: what about sysroot ?
func PkgConfigQuery(pkgspec string, cmdline []string) (PkgConfigInfo, error) {
	info := PkgConfigInfo{PkgSpec: pkgspec, Variables: make(map[string]string)}
	cmdline = append(cmdline, pkgspec)

	// modversion
	if out, err := cmd.RunOutOne(append(cmdline, "--modversion"), true); err == nil {
		info.Version = out
	} else {
		return info, err
	}

	// virtual package name
	if out, err := cmd.RunOutOne(append(cmdline, "--print-provides"), true); err == nil {
		info.Name = out
	} else {
		return info, err
	}

	// ldflags shared
	if out, err := cmd.RunOutCmd(append(cmdline, "--libs"), true); err == nil {
		info.SharedLdflags = out
	} else {
		return info, err
	}

	// ldflags static
	if out, err := cmd.RunOutCmd(append(cmdline, "--static", "--libs"), true); err == nil {
		info.StaticLdflags = out
	} else {
		return info, err
	}

	// cflags shared
	if out, err := cmd.RunOutCmd(append(cmdline, "--cflags"), true); err == nil {
		info.SharedCflags = out
	} else {
		return info, err
	}

	// cflags static
	if out, err := cmd.RunOutCmd(append(cmdline, "--static", "--cflags"), true); err == nil {
		info.StaticCflags = out
	} else {
		return info, err
	}

	if out, err := cmd.RunOutLines(append(cmdline, "--print-variables"), true); err == nil {
		for _, name := range out {
			// FIXME: need to escape
			if value, err := cmd.RunOutOne(append(cmdline, "--variable="+shellwords.Quote(name)), true); err == nil {
				info.Variables[name] = value
			}
		}
	} else {
		return info, err
	}
	return info, nil
}
