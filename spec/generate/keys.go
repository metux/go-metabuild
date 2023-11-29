package generate

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	KeyType = Key("type")

	KeyKConf    = Key("kconf")
	KeyAC       = Key("config.h")
	KeyTextfile = Key("textfile")

	KeyOutput   = Key("output")
	KeyTemplate = Key("template")
	KeyMarker   = Key("marker")
	KeyContent  = Key("content")
)
