package builder

import (
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/jobs"
)

func RunBuild(cf spec.Global) error {
	cf.BuildConf().CleanBuildDir()
	runner := jobs.NewRunner()
	if err := runner.AddJob("all-targets", MakeBuildAll(cf)); err != nil {
		util.Panicf("Build prepare error: %s", err)
	}
	if err := runner.Run(); err != nil {
		util.Panicf("Build error: %s", err)
	}
	return nil
}
