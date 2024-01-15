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
	KeyTargetObjects = Key("targets")
	KeyBuildConf     = Key("buildconf")
	KeyCache         = Key("cache")
	KeyOptions       = Key("options")
	KeyPackage       = Key("package")
	KeyVersion       = Key("version")
	KeyMaintainer    = Key("maintainer")
	KeySrcDir        = Key("srcdir")
	KeyDistro        = Key("distro")
)

// @@sys hierarchy -- internally filled fields
const (
	KeySysConfigPath    = Key("@sys::config::path")
	KeySysConfigDir     = Key("@sys::config::dir")
	KeySysConfigBase    = Key("@sys::config::base")
	KeySysConfigAbsPath = Key("@sys::config::abspath")
	KeySysConfigAbsDir  = Key("@sys::config::absdir")

	KeySysSettingsPath    = Key("@sys::settings::path")
	KeySysSettingsDir     = Key("@sys::settings::dir")
	KeySysSettingsBase    = Key("@sys::settings::base")
	KeySysSettingsAbsPath = Key("@sys::settings::abspath")
	KeySysSettingsAbsDir  = Key("@sys::settings::absdir")
)
