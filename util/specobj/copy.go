package specobj

import (
	"fmt"

	magic "github.com/metux/go-magicdict"
	"github.com/metux/go-magicdict/api"

	"github.com/metux/go-metabuild/util"
)

func (so SpecObj) copyDefaultsListFrom(src SpecObj, ent Entry, k Key) error {
	dflt, err := magic.EntryMakeList(so.Spec, k.MagicDefaults())
	util.ErrPanicf(err, "MakeList failed")

	for _, idx := range ent.Keys() {
		if e1, _ := ent.Get(idx); e1 != nil {
			if e1.IsScalar() {
				util.ErrPanicf(magic.EntryPutStr(dflt, idx, e1.String()), "adding scalar failed")
			} else if e1.IsList() {
				sublist, err3 := api.MakeList(dflt, idx)
				if sublist == nil || err != nil {
					panic(fmt.Sprintf("adding subdict failed, err=%s", err3))
				}
				so.EntrySpec(k.Append(idx)).CopyDefaultsFrom(src.EntrySpec(k.Append(idx)))
			} else if e1.IsDict() {
				subdict, err3 := api.MakeDict(dflt, idx)
				if subdict == nil || err != nil {
					panic(fmt.Sprintf("adding subdict failed, err=%s", err3))
				}
				so.EntrySpec(k.Append(idx)).CopyDefaultsFrom(src.EntrySpec(k.Append(idx)))
			}
		}
	}
	return nil
}

func (so SpecObj) CopyDefaultsFrom(src SpecObj) error {
	for _, k := range src.Spec.Keys() {
		if ent := magic.EntryGet(src.Spec, k); ent != nil {
			if ent.IsScalar() {
				so.DefaultPutStr(k, ent.String())
			} else if ent.IsList() {
				so.copyDefaultsListFrom(src, ent, k)
			} else if ent.IsDict() {
				so.EntrySpec(k).CopyDefaultsFrom(src.EntrySpec(k))
			}
		}
	}
	return nil
}
