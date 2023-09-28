package autoconf

import (
	"github.com/metux/go-metabuild/util"
)

const (
	ErrTemplateMissingMarker = util.Error("template defined, but missing marker")
	ErrCheckFailedMandatory  = util.Error("failed mandatory check")
)
