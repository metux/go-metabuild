package features

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

// per feature attributes
const (
	KeyEnabled        = Key("enabled")
	KeyPkgconfRequire = Key("pkgconf/require")
)
