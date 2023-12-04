package doc

import (
	"path/filepath"

	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
)

type ManPages struct {
	base.BaseBuilder
}

func (b ManPages) JobRun() error {

	src := b.RequiredSources()
	if len(src) != 1 {
		panic("manpage target may only have exactly one source")
	}
	srcfile := src[0]

	alias := b.EntryStrList(target.KeyManpageAlias)
	compress := b.EntryStr(target.KeyManpageCompress)

	mdir := b.InstallDir() + "man" + b.RequiredEntryStr(target.KeyManpageSection)

	// FIXME: directly write compressed file instead of copying
	switch compress {
	case "":
		// no compression
	case "gz":
		util.ErrPanicf(util.GzipCompress(srcfile, srcfile+".gz", b.InstallPerm()), "error compressing manpage")
		srcfile = srcfile + ".gz"
	default:
		panic("unsupported man compression: " + compress)
	}

	b.InstallPkgFile(srcfile, mdir, b.InstallPerm())

	for _, a := range alias {
		if compress != "" {
			a = a + "." + compress
		}
		b.InstallPkgSymlink(filepath.Base(srcfile), a, mdir)
	}

	return nil
}

func MakeManPages(o spec.TargetObject, id string) ManPages {
	return ManPages{base.BaseBuilder{o, id}}
}
