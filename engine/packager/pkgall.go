package packager

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/distro"
	"github.com/metux/go-metabuild/util/jobs"
)

type PkgAllJob struct {
	jobs.BaseJob
	Global spec.Global
}

func (job PkgAllJob) JobSub() ([]jobs.Job, error) {
	jobs := []jobs.Job{}
	bc := job.Global.BuildConf()
	for _, pkg := range bc.TargetDistro().Packages() {
		if !pkg.Skipped() {
			switch pkg.Distro.PackageFormat() {
			case distro.PkgFormatDeb:
				jobs = append(jobs, MakeDebPkgJob(bc, pkg))
			default:
				return jobs, distro.ErrPkgFormatUnsupported
			}
		}
	}
	return jobs, nil
}

func MakePkgAllJob(cf spec.Global) PkgAllJob {
	return PkgAllJob{BaseJob: jobs.BaseJob{Id: "all-packages"}, Global: cf}
}
