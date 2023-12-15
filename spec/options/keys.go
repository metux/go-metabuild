package options

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

// per option attributes
const (
	KeyEnabled         = Key("enabled")
	KeyPkgconfRequire  = Key("pkgconf/require")
	KeyAutoconfWith    = Key("autoconf/with")
	KeyAutoconfEnable  = Key("autoconf/enable")
	KeyAutoconfArgs    = Key("autoconf/args")
	KeyAutoconfDirOpts = Key("autoconf/diropts")
)
