package distro

import (
	"strings"
)

const (
	PkgFormatDEB = "deb"
	PkgFormatRPM = "rpm"
)

func DetectPkgFormat(distId string) string {
	distId = strings.ToLower(distId)

	switch distId {
	case "debian":
		return PkgFormatDEB
	case "ubuntu":
		return PkgFormatDEB
	case "sles":
		return PkgFormatRPM
	case "rhel":
		return PkgFormatRPM
	}
	return ""
}
