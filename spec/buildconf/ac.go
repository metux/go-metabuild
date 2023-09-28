package buildconf

// autoconf style config.h

import (
	"fmt"
	"strconv"
)

const (
	KeyConfigH = Key("config.h")
)

func (bc BuildConf) ACundef(sym string) {
	if sym != "" {
		bc.EntryStrListAppend(KeyConfigH, fmt.Sprintf("#undef CONFIG_%s", sym))
	}
}

func (bc BuildConf) ACraw(sym string, val string) {
	if sym != "" {
		bc.EntryStrListAppend(KeyConfigH, fmt.Sprintf("#define CONFIG_%s %s", sym, val))
	}
}

func (bc BuildConf) ACBool(sym string, val bool) {
	if val {
		bc.ACraw(sym, "1")
	} else {
		bc.ACundef(sym)
	}
}

func (bc BuildConf) ACStr(sym string, val string) {
	bc.ACraw(sym, strconv.Quote(val))
}

func (bc BuildConf) ACLines() []string {
	return bc.EntryStrList(KeyConfigH)
}
