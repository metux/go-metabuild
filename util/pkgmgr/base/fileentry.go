package base

import (
	"strings"
)

type PkgFileEntry struct {
	Name    string
	Package string
	PkgName string
	PkgArch string
}

func NewPkgFileEntry(pkg string, fn string) PkgFileEntry {
	psplit := strings.Split(pkg, ":")
	if len(psplit) == 1 {
		return PkgFileEntry{fn, pkg, pkg, ""}
	}
	return PkgFileEntry{fn, pkg, psplit[0], psplit[1]}
}
