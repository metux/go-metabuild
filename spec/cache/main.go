package cache

import (
	"github.com/metux/go-metabuild/util/specobj"
)

type Cache struct {
	specobj.SpecObj
}

// FIXME: should we take a subspec ?
func (c Cache) CheckGet(h string) (bool, bool) {
	k := KeyChecks.AppendStr(h)
	if c.EntryBoolDef(k.Append(KeyCached), false) {
		return true, c.EntryBoolDef(k.Append(KeyResult), false)
	}
	return false, false
}

func NewCache(so specobj.SpecObj) Cache {
	return Cache{so}
}
