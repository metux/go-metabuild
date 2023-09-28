package packager

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util/jobs"
)

func RunPackaging(cf spec.Global) error {
	runner := jobs.NewRunner()
	if err := runner.AddJob("all-packages", MakePkgAllJob(cf)); err != nil {
		return err
	}
	return runner.Run()
}
