package distro

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	DistTypeDebian = "debian"
	DistTypeRHEL   = "rhel"
	DistTypeSLES   = "sles"
)

// attributes of per-distro blocks
const (
	KeyPackageFormat = Key("pkg-format")
	KeyPackages      = Key("packages")
)

// supported packaging formats (KeyPackageFormat)
const (
	PkgFormatDeb = "deb"
	PkgFormatRpm = "rpm"
)
