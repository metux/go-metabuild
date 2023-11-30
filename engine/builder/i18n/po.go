package i18n

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
)

type I18nPo struct {
	base.BaseBuilder
}

const (
	SuffixPo = ".po"
	SuffixMo = ".mo"
)

func (b I18nPo) JobRun() error {
	srcdir := b.RequiredEntryStr(target.KeySourceDir) + "/"
	catdir := b.RequiredEntryStr(target.KeyI18nCategory) + "/"
	perm := b.InstallPerm()
	installdir := b.InstallDir() + "/"

	fn := b.RequiredEntryStr(target.KeyI18nDomain) + SuffixMo

	// FIXME: probe this, handle endianess ?
	prog := []string{"msgfmt"}

	for _, l := range b.FeaturedStrList(target.KeyI18nLinguas) {
		subdir := l + "/" + catdir
		outfile := b.BuildConf.BuildTempDir("po/"+subdir) + "/" + fn
		infile := srcdir + l + SuffixPo

		b.ExecAbort(append(prog, "-o", outfile, infile), "")
		b.InstallPkgFile(outfile, installdir+subdir, perm)
	}
	return nil
}

func MakeI18nPo(o spec.TargetObject, id string) I18nPo {
	return I18nPo{base.BaseBuilder{o, id}}
}
