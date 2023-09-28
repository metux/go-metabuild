package specobj

import (
	magic "github.com/metux/go-magicdict"
)

type Key = magic.Key
type Entry = magic.Entry

type SpecObj struct {
	Spec Entry
	Err  error
}

func (so SpecObj) MyKey() Key {
	return Key(so.EntryStr("@@KEY"))
}

func NewSpecObj(ent magic.Entry) SpecObj {
	return SpecObj{ent, nil}
}

func NewSpecObjErr(ent magic.Entry, err error) SpecObj {
	return SpecObj{ent, err}
}
