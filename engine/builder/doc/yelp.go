package doc

import (
	"path/filepath"

	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/fileutil"
)

type Yelp struct {
	base.BaseBuilder
}

func (b Yelp) installFiles(files []string) {
	installdir := b.InstallDir()
	fmode := b.InstallPerm()
	for _, e := range files {
		b.InstallPkgFile(e, installdir, fmode)
	}
}

func (b Yelp) instDir(dir string, suffix string) error {
	entries, err := fileutil.ListDir(dir, suffix)
	if err != nil {
		return err
	}

	b.installFiles(entries)
	return nil
}

func (b Yelp) JobRun() error {
	src := b.RequiredSourceAbs()
	format := b.RequiredEntryStr(target.KeyOutputFormat)
	includes := b.EntryPathList(target.KeyIncludeDir)
	outdir := b.BuildConf.BuildTempDir("yelp/" + b.MyId() + "/" + format)

	c := append(b.BuilderCmd(), format, "-o", outdir)
	for _, d := range includes {
		c = append(c, "-p", d)
	}
	c = append(c, src)

	b.ExecAbort(c, filepath.Dir(src))

	if b.WantInstall() {
		if err := b.instDir(outdir, ""); err != nil {
			return err
		}
		for _, d := range includes {
			if err := b.instDir(d, ".png"); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b Yelp) JobPrepare(id string) error {
	return nil
}

func MakeYelp(o spec.TargetObject, id string) Yelp {
	return Yelp{base.BaseBuilder{o, id}}
}
