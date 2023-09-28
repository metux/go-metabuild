package packager

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/fileutil"
	"github.com/metux/go-metabuild/util/jobs"
	"github.com/metux/go-metabuild/util/pkgmgr"
)

type BasePkgJob struct {
	jobs.BaseJob
	Packager  pkgmgr.PackageManager
	BuildConf spec.BuildConf
	Pkg       spec.DistPkg
}

func (job BasePkgJob) MetaDir() string {
	return job.BuildConf.BuildDistPkgMetaDir(job.Pkg.Name())
}

func (job BasePkgJob) FindLocalLib(fn string, arch string) string {
	dist := job.Pkg.Distro
	for _, pId := range dist.PackageIds() {
		pname := dist.PackageNameTrans(string(pId))
		pkgroot := job.BuildConf.BuildDistPkgRootDir(pname, "")
		flist := fileutil.FindFile(pkgroot, fn)

		//		// FIXME: need to check arch
		//		for _,a := range flist {
		//			return pname
		//		}

		if len(flist) > 0 {
			return pname
		}
	}
	return ""
}

// find package that provides given shared object
func (job BasePkgJob) FindSoPkg(fn string, arch string) string {
	// try to find it locally (in our own package
	if localpkg := job.FindLocalLib(fn, arch); localpkg != "" {
		return localpkg
	}

	// query the package manager
	for _, x := range job.Packager.SearchInstalledFile(fn) {
		if job.Packager.MatchElfArch(x.PkgArch, arch) {
			return x.Package
		}
	}
	return ""
}

func (job BasePkgJob) FindPcPkg(pcname string, version string) string {
	for _, x := range job.Packager.SearchInstalledFile(pcname + ".pc") {
		return x.PkgName + " (>=" + version + ")"
	}
	return ""
}

func (job BasePkgJob) ScanSoDep() ([]string, error) {
	deps := []string{}

	lines, err := fileutil.ReadLinesGlobUniq(job.MetaDir() + "/*.sodep")
	if err != nil {
		return deps, err
	}

	for _, l1 := range lines {
		entry := strings.Fields(l1)
		if len(entry) < 2 {
			continue
		}
		pkgname := job.FindSoPkg(entry[0], entry[1])
		if pkgname == "" {
			return deps, fmt.Errorf("cant find package providing %s %s", entry[0], entry[1])
		}
		deps = append(deps, pkgname)
	}

	return util.Uniq(deps), nil
}

func (job BasePkgJob) ScanPcDep() ([]string, error) {
	deps := []string{}

	lines, err := fileutil.ReadLinesGlobUniq(job.MetaDir() + "/*.pcdep")
	if err != nil {
		return deps, err
	}

	for _, l1 := range lines {
		entry := strings.Fields(l1)
		if len(entry) < 2 {
			continue
		}
		pkgname := job.FindPcPkg(entry[0], entry[1])
		if pkgname == "" {
			return deps, fmt.Errorf("cant find package providing %s", l1)
		}
		deps = append(deps, pkgname)
	}

	return util.Uniq(deps), nil
}

func (job BasePkgJob) ScanAutoDep() ([]string, error) {
	sodep, err := job.ScanSoDep()
	if err != nil {
		return sodep, err
	}
	pcdep, err := job.ScanPcDep()
	if err != nil {
		return pcdep, err
	}

	deps := util.Uniq(append(sodep, pcdep...))
	return deps, nil
}

func (job BasePkgJob) ScanTriggers() []string {
	lines, _ := fileutil.ReadLinesGlobUniq(job.MetaDir() + "/*.trigger")
	return lines
}

func (job BasePkgJob) DoPackage() error {
	return job.Packager.Build(
		job.BuildConf.BuildDistPkgRootDir(job.Pkg.Name(), ""),
		job.BuildConf.BuildDistDir(""))
}

func (job BasePkgJob) DataRoot() string {
	return filepath.Clean(job.BuildConf.BuildDistPkgRootDir(job.Pkg.Name(), "")) + "/"
}

func MakeBasePkgJob(packager pkgmgr.PackageManager, buildconf spec.BuildConf, pkg spec.DistPkg) BasePkgJob {
	return BasePkgJob{
		BaseJob:   jobs.MakeBaseJob("package-" + pkg.Id()),
		Packager:  packager,
		BuildConf: buildconf,
		Pkg:       pkg}
}
