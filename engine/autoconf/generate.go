package autoconf

import (
	"log"
	"strings"

	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/generate"
	"github.com/metux/go-metabuild/util/fileutil"
)

func RunGenerates(cf spec.Global) error {
	for _, g := range cf.GetGenerates() {
		if err := runGen(g); err != nil {
			return err
		}
	}
	return nil
}

func runGenLines(g spec.Generate) (string, error) {
	switch spec.Key(g.Type()) {

	// generate autoconf style config.h
	case generate.KeyAC:
		return strings.Join(g.BuildConf.ACLines(), "\n"), nil

	// generate kconfig style .config
	case generate.KeyKConf:
		return strings.Join(g.BuildConf.KConfLines(), "\n"), nil
	}

	return "", generate.ErrUnsupportedGenerate
}

func runGen(g spec.Generate) error {
	fn := g.OutputFile()

	lines, err := runGenLines(g)
	if err != nil {
		return err
	}

	if tmpl := g.TemplateFile(); tmpl != "" {
		if marker := g.Marker(); marker != "" {
			log.Printf("generating %s from template %s marker %s\n", fn, tmpl, marker)
			return fileutil.WriteTemplate(tmpl, fn, marker, lines)
		}
		return ErrTemplateMissingMarker
	}

	log.Println("generating:", fn)
	return fileutil.WriteText(fn, lines)
}
