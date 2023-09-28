package buildconf

// kernel/kconf style .config

import (
	"fmt"
)

const (
	KeyKConf = Key("kconf")
)

func (bc BuildConf) KConfUndef(sym string) {
	if sym != "" {
		bc.EntryStrListAppend(KeyKConf, fmt.Sprintf("# CONFIG_%s is not defined", sym))
	}
}

func (bc BuildConf) KConfRaw(sym string, val string) {
	if sym != "" {
		bc.EntryStrListAppend(KeyKConf, fmt.Sprintf("CONFIG_%s=%s", sym, val))
	}
}

func (bc BuildConf) KConfBool(sym string, val bool) {
	if val {
		bc.KConfRaw(sym, "y")
	} else {
		bc.KConfUndef(sym)
	}
}

func (bc BuildConf) KConfStr(sym string, val string) {
	if val != "" {
		bc.KConfRaw(sym, val)
	} else {
		bc.KConfUndef(sym)
	}
}

func (bc BuildConf) KConfLines() []string {
	return bc.EntryStrList(KeyKConf)
}
