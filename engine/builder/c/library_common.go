package c

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/util/compiler"
	"github.com/metux/go-metabuild/util/jobs"
)

// base for library components
type CommonCBuilder struct {
	base.BaseBuilder
	Compiler compiler.CompilerInfo
	Parent   *BaseCBuilder
}

// the child jobs must depend on the parent's deps, so all imported libs
// are already built
//
// FIXME: we could achieve better load balancing when depending on just specific
// child jobs of the parent's Deps (eg. static lib doesn't need the static
// deps already built), but that would become quite complex to implement.
func (b CommonCBuilder) JobDepends() []jobs.JobId {
	return b.Parent.JobDepends()
}
