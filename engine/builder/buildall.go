package builder

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util/jobs"
)

type BuildAll struct {
	spec.Global
	jobs.BaseJob
}

func (b BuildAll) JobSub() ([]jobs.Job, error) {
	jobs := []jobs.Job{}
	for id, t := range b.GetTargetObjects() {
		if b, err := CreateBuilder(t, id); err == nil {
			jobs = append(jobs, b)
		} else {
			return jobs, err
		}
	}
	return jobs, nil
}

func MakeBuildAll(g spec.Global) BuildAll {
	return BuildAll{Global: g, BaseJob: jobs.BaseJob{Id: "all-targets"}}
}
