package c

import (
	"os"

	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
)

type BuilderCLibraryDevlink struct {
	BaseCBuilder
}

func (b BuilderCLibraryDevlink) JobRun() error {

	dest := b.RequiredEntryStr(target.KeyLinkTarget)
	outname := b.OutputFile()

	os.Remove(outname)

	util.ErrPanicf(
		os.Symlink(dest, outname),
		"Failed creating symlink %s", outname)

	if b.WantInstall() {
		b.InstallPkgSymlink(dest, outname, b.InstallDir())
	}
	return nil
}
