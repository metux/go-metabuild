package base

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/fileutil"
	"github.com/metux/go-metabuild/util/jobs"
)

type Key = spec.Key

type BaseBuilder struct {
	spec.TargetObject
	Id string
}

func (b *BaseBuilder) BaseInit(id string) error {
	b.Id = id
	return nil
}

func (b BaseBuilder) Logf(format string, v ...any) {
	log.Printf("[Job "+b.Id+"] "+format, v...)
}

func (b BaseBuilder) PkgPath(pkg string, dst string) string {
	return b.BuildConf.BuildDistPkgRootDir(b.BuildConf.PkgNameTrans(pkg), dst)
}

func (b BaseBuilder) InstallPkgFileAuto() bool {
	if b.WantInstall() {
		b.InstallPkgFile(b.OutputFile(), b.InstallDir(), b.InstallPerm())
		return true
	}
	return false
}

func (b BaseBuilder) InstallPkgFile(src string, dst string, mode os.FileMode) {
	pkg := b.PackageId("")
	if mode == 0 {
		panic("config error: file mode is 0")
	}
	targetdir := b.PkgPath(pkg, dst)
	b.Logf("Installing [%s]: src=%s dst=%s fmode=%s dir=%s\n", pkg, src, dst, mode, targetdir)
	if err := fileutil.CopyFile(src, targetdir+"/"+filepath.Base(src), mode); err != nil {
		panic(fmt.Sprintf("failed copy file: %s", err))
	}
}

func (b BaseBuilder) InstallPkgFileCompressed(src string, dst string, mode os.FileMode, compress string) {
	pkg := b.PackageId("")
	if mode == 0 {
		panic("config error: file mode is 0")
	}
	targetdir := b.PkgPath(pkg, dst)
	b.Logf("Installing [%s]: compress=%s src=%s dst=%s fmode=%s dir=%s\n", compress, pkg, src, dst, mode, targetdir)

	switch compress {
	case "":
		b.InstallPkgFile(src, dst, mode)
	case "gz":
		util.ErrPanicf(util.GzipCompress(src, targetdir+"/"+filepath.Base(src)+".gz", mode), "compression failed")
	}
}

func (b BaseBuilder) InstallPkgSymlink(src string, dst string, dir string) {
	pkg := b.PackageId("")
	targetdir := b.PkgPath(pkg, dir)
	b.Logf("Linking [%s]: src=%s dst=%s dir=%s targetdir=%s\n", pkg, src, dst, dir, targetdir)
	os.Remove(targetdir + "/" + dst)
	if err := os.Symlink(src, targetdir+"/"+dst); err != nil {
		panic(fmt.Sprintf("failed creating symlink: %s", err))
	}
}

func (b BaseBuilder) WritePkgMeta(id string, data string) {
	pkg := b.PackageId("")
	fn := b.BuildConf.BuildDistPkgMetaDir(b.BuildConf.PkgNameTrans(pkg)) + "/" + id
	b.Logf("Writing [%s]: fn=%s\n", pkg, fn)
	if err := fileutil.WriteText(fn, data); err != nil {
		panic(fmt.Sprintf("WritePkgMeta failed: %s", err))
	}
}

func (b BaseBuilder) JobId() jobs.JobId {
	return b.Id
}

func (b BaseBuilder) JobSub() ([]jobs.Job, error) {
	return []jobs.Job{}, nil
}

func (b BaseBuilder) JobRun() error {
	return nil
}

func (b BaseBuilder) JobPrepare(id jobs.JobId) error {
	b.LoadTargetDefaults()
	return nil
}

func (b BaseBuilder) JobDepends() []jobs.JobId {
	return []jobs.JobId{}
}

func (b BaseBuilder) TempDir() string {
	return b.BuildConf.BuildTempDir("jobs/" + b.Id)
}

func (b BaseBuilder) OutputFile() string {
	s := b.RequiredEntryStr(target.KeyFile)
	os.MkdirAll(filepath.Dir(s), 0755)
	return s
}

func NewBaseBuilder(o spec.TargetObject, id string) BaseBuilder {
	return BaseBuilder{o, id}
}
