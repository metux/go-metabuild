package packager

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util/fileutil"
	"github.com/metux/go-metabuild/util/pkgmgr"
)

type DebPkgJob struct {
	BasePkgJob
}

func (job DebPkgJob) JobRun() error {
	dataroot := job.DataRoot()

	trig := job.ScanTriggers()
	if len(trig) != 0 {
		// FIXME: move this to packager util
		if err := fileutil.WriteFileLines(dataroot+"/DEBIAN/triggers", trig); err != nil {
			return err
		}
	}

	autodeps, err := job.ScanAutoDep()
	if err != nil {
		return err
	}

	if err = job.Packager.WriteControlFile(dataroot, job.Pkg.ControlInfo(autodeps)); err != nil {
		return err
	}

	return job.DoPackage()
}

func MakeDebPkgJob(buildconf spec.BuildConf, pkg spec.DistPkg) DebPkgJob {
	// FIXME: need to load platform settings
	packager := pkgmgr.NewDpkg([]string{}, []string{}, "")
	return DebPkgJob{MakeBasePkgJob(packager, buildconf, pkg)}
}
