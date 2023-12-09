package composite

import (
	"github.com/metux/go-metabuild/engine/builder/c"
	"github.com/metux/go-metabuild/engine/builder/gen"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util/jobs"
)

type CompositeGlibResource struct {
	c.BaseCBuilder
}

func (b CompositeGlibResource) jobGen() jobs.Job {
	builderGen := gen.MakeGlibResource(b.SubTarget("generate"), b.JobId()+"/generate")
	builderGen.LoadTargetDefaults()
	return builderGen
}

func (b CompositeGlibResource) jobLib() jobs.Job {
	builderLib := c.MakeCLibrary(b.SubTarget("library"), b.JobId()+"/library")
	builderLib.LoadTargetDefaults()
	return builderLib
}

func (b CompositeGlibResource) JobSub() ([]jobs.Job, error) {
	jobs := []jobs.Job{b.jobGen(), b.jobLib()}

	// needed a 2nd time, so we can overwrite sub's template by ours
	b.LoadTargetDefaults()

	return jobs, nil
}

func MakeGlibResource(o spec.TargetObject, id string) CompositeGlibResource {
	return CompositeGlibResource{BaseCBuilder: c.MakeBaseCBuilder(o, id)}
}
