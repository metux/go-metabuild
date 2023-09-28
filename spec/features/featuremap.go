package features

import (
	"github.com/metux/go-metabuild/util/specobj"
)

type FeatureMap struct {
	specobj.SpecObj
}

func (fm FeatureMap) IDs() []Key {
	return fm.Keys()
}

func (fm FeatureMap) All() []Feature {
	f := []Feature{}
	for _, k := range fm.Keys() {
		f = append(f, Feature{fm.EntrySpec(k), k})
	}
	return f
}

func (fm FeatureMap) Map() map[Key]string {
	m := make(map[Key]string)
	for _, k := range fm.Keys() {
		m[k] = fm.EntryStr(k.Append(KeyEnabled))
	}
	return m
}

func (fm FeatureMap) Init() {
	for _, k := range fm.Keys() {
		fm.DefaultPutStr(k.Append(KeyEnabled), "${@@PARENT::default}")
	}
}
