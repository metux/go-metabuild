package c

import (
	"github.com/metux/go-metabuild/engine/builder/base"
)

// base for library components
type CommonCBuilder struct {
	base.BaseBuilder
	Parent *BaseCBuilder
}
