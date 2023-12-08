package composite

import (
	"github.com/metux/go-metabuild/engine/builder/c"
	"github.com/metux/go-metabuild/engine/builder/gen"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/util/jobs"
)

type CompositeGlibMarshal struct {
	c.BaseCBuilder
}

func (b CompositeGlibMarshal) jobGen() jobs.Job {
	builderGen := gen.MakeGlibMarshal(b.SubTarget("generate"), b.JobId()+"/generate")
	builderGen.LoadTargetDefaults()
	return builderGen
}

func (b CompositeGlibMarshal) jobLib() jobs.Job {
	builderLib := c.MakeCLibrary(b.SubTarget("library"), b.JobId()+"/library")
	builderLib.LoadTargetDefaults()
	return builderLib
}

func (b CompositeGlibMarshal) JobSub() ([]jobs.Job, error) {
	jobs := []jobs.Job{b.jobGen(), b.jobLib()}

	// needed a 2nd time, so we can overwrite sub's template by ours
	b.LoadTargetDefaults()

	return jobs, nil
}

func MakeGlibMarshal(o spec.TargetObject, id string) CompositeGlibMarshal {
	return CompositeGlibMarshal{BaseCBuilder: c.MakeBaseCBuilder(o, id)}
}
