package git

import (
	"github.com/metux/go-metabuild/util/cmd"
)

func Describe(tags bool, all bool, exclude []string, match []string) (string, error) {
	c := append(gitCmd, "describe")

	if tags {
		c = append(c, "--tags")
	}

	if all {
		c = append(c, "--all")
	}

	for _, ex := range exclude {
		c = append(c, "--exclude", ex)
	}

	for _, ex := range match {
		c = append(c, "--match", ex)
	}

	return cmd.RunOutOne(c, true)
}

func PkgVersion(tags bool, all bool, exclude []string, match []string) (string, error) {
	ver, err := Describe(tags, all, exclude, match)
	if err == nil {
		if len(ver) > 1 && ver[0] == 'v' {
			ver = ver[1:]
		}
	}
	return ver, err
}
