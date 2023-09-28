package cache

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	KeyChecks = Key("checks")

	KeyCached = Key("cached")
	KeyResult = Key("result")
)
