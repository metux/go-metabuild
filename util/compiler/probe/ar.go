package probe

import (
	"os"

	"github.com/metux/go-metabuild/util/compiler/base"
)

func probeArCmd(ci *base.CompilerInfo) string {
	// is this the cross compiler ?
	if ci.CrossForHost {
		if e := os.Getenv("AR"); e != "" {
			return e
		}
		return ci.CrossPrefix + "ar"
	}

	// is this the cross build system ?
	if ci.CrossForBuild {
		if e := os.Getenv("HOST_AR"); e != "" {
			return e
		}
		return "ar"
	}

	// normal compile
	if e := os.Getenv("AR"); e != "" {
		return e
	}
	return "ar"
}
