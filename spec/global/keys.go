package global

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	KeyConfigureChecks   = Key("configure::checks")
	KeyConfigureGenerate = Key("configure::generate")

	KeyCheckedCDefines = Key("buildrun::compile-c::symbols")
)

// subkeys of global spec
const (
	KeyTargetObjects = Key("objects")
	KeyBuildConf     = Key("buildconf")
	KeyCache         = Key("cache")
	KeyFeatures      = Key("features")
	KeyPackage       = Key("package")
	KeyVersion       = Key("version")
	KeyMaintainer    = Key("maintainer")
	KeySrcDir        = Key("srcdir")
	KeyDistro        = Key("distro")
)
