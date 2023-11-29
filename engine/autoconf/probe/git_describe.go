package probe

import (
	"strings"

	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/util/git"
)

type GitDescribe struct {
	ProbeBase
}

func (p GitDescribe) Probe() error {
	if ver, err := git.PkgVersion(
		p.EntryBoolDef("tags", false),
		p.EntryBoolDef("all", false),
		p.EntryStrList("exclude"),
		p.EntryStrList("match")); err == nil {
		p.Logf("detected version from git: \"%s\"\n", ver)
		p.BuildConf.EntryPutStr(buildconf.KeyVersion, ver)
	} else {
		return err
	}

	if authors, err := git.Authors(); err == nil {
		p.BuildConf.EntryPutStr(buildconf.KeyAuthors, strings.Join(authors, "\n"))
	} else {
		return err
	}

	return nil
}

func MakeGitDescribe(chk Check) ProbeInterface {
	return GitDescribe{MakeProbeBase(chk)}
}
