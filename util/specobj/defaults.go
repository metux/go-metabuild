package specobj

import (
	magic "github.com/metux/go-magicdict"
)

func (so SpecObj) DefaultPutStr(k Key, v string) error {
	return magic.DefaultPutStr(so.Spec, k, v)
}

func (so SpecObj) DefaultStr(k Key) string {
	return magic.DefaultStr(so.Spec, k)
}

// Note: this function needs key+value pairs, thus parameter count *must* be even
// passing an odd number of parameters will cause panic
func (so SpecObj) DefaultPutStrs(strs ...string) {
	l := len(strs) / 2

	for x := 0; x < l; x++ {
		so.DefaultPutStr(Key(strs[x*2]), strs[x*2+1])
	}
}

func (so SpecObj) DefaultPutStrMap(m map[Key]string) {
	for k, v := range m {
		so.DefaultPutStr(k, v)
	}
}
