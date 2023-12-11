package features

import (
	"github.com/metux/go-metabuild/util/specobj"
)

type Feature struct {
	specobj.SpecObj
	Id Key
}

func (f Feature) Value() string {
	return f.EntryStr(KeyEnabled)
}

func (f Feature) ValueYN() string {
	switch v := f.Value(); v {
	case "true", "yes", "1":
		return "y"
	case "false", "no", "0", "":
		return "n"
	default:
		return v
	}
}

func (f Feature) IsOn() bool {
	switch v := f.Value(); v {
	case "y", "true", "yes", "1":
		return true
	case "n", "false", "no", "0", "":
		return false
	}
	return false
}

func (f Feature) Set(v string) {
	f.EntryPutStr(KeyEnabled, v)
}

func (f Feature) SetOn() {
	f.Set("on")
}

func (f Feature) SetOff() {
	f.Set("off")
}

func (f Feature) FlagsOn() specobj.SpecObj {
	return f.EntrySpec(Key("set@" + f.Value()))
}

func (f Feature) PkgconfRequire() []string {
	return f.EntryStrList(KeyPkgconfRequire)
}
