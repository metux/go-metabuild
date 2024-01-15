package options

import (
	"github.com/metux/go-metabuild/util/specobj"
)

type OptionMap struct {
	specobj.SpecObj
}

func (fm OptionMap) IDs() []Key {
	return fm.Keys()
}

func (fm OptionMap) All() []Feature {
	f := []Feature{}
	for _, k := range fm.Keys() {
		f = append(f, fm.Get(k))
	}
	return f
}

func (fm OptionMap) Get(k Key) Feature {
	return Feature{fm.EntrySpec(k), k}
}

func (fm OptionMap) Map() map[Key]string {
	m := make(map[Key]string)
	for _, k := range fm.Keys() {
		m[k] = fm.EntryStr(k.Append(KeyEnabled))
	}
	return m
}

func (fm OptionMap) Init() {
	for _, k := range fm.Keys() {
		fm.DefaultPutStr(k.Append(KeyEnabled), "${@@PARENT::default}")
	}
}
