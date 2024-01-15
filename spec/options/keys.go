package options

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

// per option attributes
const (
	KeyEnabled        = Key("enabled")
	KeyPkgconfRequire = Key("pkgconf/require")
)
