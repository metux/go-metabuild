package generate

import (
	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/util/specobj"
)

type Generate struct {
	specobj.SpecObj
	BuildConf buildconf.BuildConf
}

type GenerateList = []Generate

func (gen Generate) Init() error {
	return gen.detectType()
}

func (gen Generate) String() string {
	return gen.Type() + " => " + gen.OutputFile()
}

func (gen Generate) OutputFile() string {
	return gen.EntryStr(KeyOutput)
}

func (gen Generate) TemplateFile() string {
	return gen.EntryStr(KeyTemplate)
}

func (gen Generate) Marker() string {
	return gen.EntryStr(KeyMarker)
}

func (gen Generate) Type() string {
	return gen.EntryStr(KeyType)
}

func (gen Generate) detectType() error {
	if gen.Type() != "" {
		return nil
	}

	types := []Key{
		KeyKConf,
		KeyAC,
		KeyTextfile,
	}

	for _, x := range types {
		if s := gen.EntryStrList(x); len(s) > 0 {
			gen.EntryPutStr(KeyType, string(x))
			gen.EntryPutStr(KeyOutput, s[0])
			return nil
		}
	}

	return ErrUnsupportedGenerate
}
