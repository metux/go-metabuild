package specobj

import (
	magic "github.com/metux/go-magicdict"
)

func (so SpecObj) EntryBoolDef(k Key, dflt bool) bool {
	return magic.EntryBoolDef(so.Spec, k, dflt)
}

func (so SpecObj) EntryPutBool(k Key, v bool) error {
	return magic.EntryPutBool(so.Spec, k, v)
}

func (so SpecObj) EntryIntDef(k Key, dflt int) int {
	return magic.EntryIntDef(so.Spec, k, dflt)
}

func (so SpecObj) EntryPutInt(k Key, v int) error {
	return magic.EntryPutInt(so.Spec, k, v)
}

func (so SpecObj) EntryStr(k Key) string {
	return magic.EntryStr(so.Spec, k)
}

func (so SpecObj) EntryStrList(k Key) []string {
	return magic.EntryStrList(so.Spec, k)
}

func (so SpecObj) EntryPutStrList(k Key, s []string) error {
	return magic.EntryPutStrList(so.Spec, k, s)
}

func (so SpecObj) EntryPutStr(k Key, v string) error {
	return magic.EntryPutStr(so.Spec, k, v)
}

func (so SpecObj) EntryPutStrMap(k Key, v map[string]string) error {
	return magic.EntryPutStrMap(so.Spec, k, v)
}

func (so SpecObj) EntryDelete(k Key) {
	magic.EntryDelete(so.Spec, k)
}

func (so SpecObj) EntryStrListAppend(k Key, s string) {
	magic.EntryStrListAppend(so.Spec, k, s)
}

func (so SpecObj) EntryStrListAppendList(k Key, s []string) {
	for _, x := range s {
		magic.EntryStrListAppend(so.Spec, k, x)
	}
}

func (so SpecObj) EntryKeys(k Key) []Key {
	return magic.EntryKeys(so.Spec, k)
}

func (so SpecObj) EntryElems(k Key) []Entry {
	return magic.EntryElems(so.Spec, k)
}

func (so SpecObj) EntryStrMap(k Key) map[Key]string {
	return magic.EntryStrMap(so.Spec, k)
}

func (so SpecObj) EntryMap(k Key) map[Key]Entry {
	return magic.EntryMap(so.Spec, k)
}

func (so SpecObj) EntrySpec(k Key) SpecObj {
	if sub, err := magic.EntryMakeDict(so.Spec, k); err != nil {
		return NewSpecObjErr(sub, err)
	} else {
		return NewSpecObj(sub)
	}
}

func (so SpecObj) EntrySpecMap(k Key) map[Key]SpecObj {
	m := make(map[Key]SpecObj)
	for k, v := range so.EntryMap(k) {
		m[k] = NewSpecObj(v)
	}
	return m
}

func (so SpecObj) EntrySpecList(k Key) []SpecObj {
	elems := so.EntryElems(k)
	ret := make([]SpecObj, len(elems))
	for idx, ent := range elems {
		ret[idx] = NewSpecObj(ent)
	}
	return ret
}

func (so SpecObj) RequiredEntryStr(k Key) string {
	return magic.RequiredEntryStr(so.Spec, k)
}

func SpecXformList[K ~string, V any](so SpecObj, k Key, proc func(K) V) []V {
	names := so.EntryStrList(k)
	data := make([]V, len(names), len(names))
	for idx, name := range names {
		data[idx] = proc(K(name))
	}
	return data
}

func SpecXformSpecList[V any](so SpecObj, k Key, proc func(SpecObj) V) []V {
	vals := so.EntrySpecList(k)
	data := make([]V, len(vals), len(vals))
	for idx, val := range vals {
		data[idx] = proc(val)
	}
	return data
}
