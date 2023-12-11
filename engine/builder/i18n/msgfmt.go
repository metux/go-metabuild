package i18n

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/fileutil"
)

type I18nMsgfmt struct {
	base.BaseBuilder
}

// FIXME: scan for linguas
func (b I18nMsgfmt) JobRun() error {
	linguas := b.RequiredEntryStrList(target.KeyI18nLinguas)
	podir := b.RequiredEntryStr(target.KeyI18nPoDir)
	subdir := b.EntryStr(target.KeySourceDir)
	inSuffix := b.RequiredEntryStr(target.KeySourceSuffix)
	outSuffix := b.RequiredEntryStr(target.KeyOutputSuffix)
	installdir := b.InstallDir()
	typ := "--" + b.RequiredEntryStr(target.KeyI18nType)

	perm := b.InstallPerm()

	// write linguas file
	util.ErrPanicf(fileutil.WriteFileLines(podir+"/LINGUAS", linguas), "failed writing LINGUAS file")

	tmpdir := b.BuildConf.BuildTempDir(string(b.Type()))

	for _, item := range b.FeaturedStrList(target.KeySource) {
		infile := fileutil.MkPath(subdir, item+inSuffix)
		tmpfile := fileutil.MkPath(tmpdir, item+inSuffix)
		outfile := fileutil.MkPath(tmpdir, item+outSuffix)

		// need to fix intltool specific syntax
		lines := fileutil.ReadLines(infile)
		for idx, l := range lines {
			if len(l) > 0 && l[0] == '_' {
				lines[idx] = l[1:]
			}
		}

		fileutil.WriteFileLines(tmpfile, lines)

		b.ExecAbort(append(b.BuilderCmd(), typ, "-d", podir, "--template", tmpfile, "-o", outfile), "")

		b.InstallPkgFile(outfile, installdir, perm)
	}
	return nil
}

func MakeI18nMsgfmt(o spec.TargetObject, id string) I18nMsgfmt {
	return I18nMsgfmt{base.BaseBuilder{o, id}}
}
