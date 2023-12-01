package i18n

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/fileutil"
)

type I18nPo struct {
	base.BaseBuilder
}

const (
	SuffixPo = ".po"
	SuffixMo = ".mo"
)

func (b I18nPo) JobRun() error {
	srcdir := b.RequiredEntryStr(target.KeySourceDir)
	catdir := b.RequiredEntryStr(target.KeyI18nCategory)
	perm := b.InstallPerm()
	installdir := b.InstallDir()
	fn := b.RequiredEntryStr(target.KeyI18nDomain) + SuffixMo
	prog := b.BuilderCmd()

	for _, l := range b.FeaturedStrList(target.KeyI18nLinguas) {
		subdir := fileutil.MkPath(l, catdir)
		outfile := fileutil.MkPath(b.BuildConf.BuildTempDir("po/"+subdir), fn)
		infile := fileutil.MkPath(srcdir, l+SuffixPo)

		b.ExecAbort(append(prog, "-o", outfile, infile), "")
		b.InstallPkgFile(outfile, fileutil.MkPath(installdir, subdir), perm)
	}
	return nil
}

func MakeI18nPo(o spec.TargetObject, id string) I18nPo {
	return I18nPo{base.BaseBuilder{o, id}}
}
