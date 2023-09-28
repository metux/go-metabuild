package specobj

// implementation of proxy to the magicdict.Entry interface

import (
	magic "github.com/metux/go-magicdict"
)

func (so SpecObj) Get(k Key) (Entry, error) {
	if so.Spec == nil {
		return nil, magic.ErrNilInterface
	}
	return so.Spec.Get(k)
}

func (so SpecObj) Put(k Key, value Entry) error {
	if so.Spec == nil {
		return magic.ErrNilInterface
	}
	return so.Spec.Put(k, value)
}

func (so SpecObj) Keys() magic.KeyList {
	if so.Spec == nil {
		return magic.KeyList{}
	}
	return so.Spec.Keys()
}

func (so SpecObj) Elems() magic.EntryList {
	if so.Spec == nil {
		return magic.EntryList{}
	}
	return so.Spec.Elems()
}

func (so SpecObj) Empty() bool {
	if so.Spec == nil {
		return true
	}
	return so.Spec.Empty()
}

func (so SpecObj) String() string {
	if so.Spec == nil {
		return ""
	}
	return so.Spec.String()
}

func (so SpecObj) IsConst() bool {
	if so.Spec == nil {
		return false
	}
	return so.Spec.IsConst()
}

func (so SpecObj) MayMergeDefaults() bool {
	if so.Spec == nil {
		return false
	}
	return so.Spec.MayMergeDefaults()
}

func (so SpecObj) IsScalar() bool {
	if so.Spec == nil {
		return false
	}
	return so.Spec.IsScalar()
}
