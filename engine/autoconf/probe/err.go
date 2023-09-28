package probe

import (
	"github.com/metux/go-metabuild/util"
)

const (
	ErrCompileFailed = util.Error("compile failed")
	ErrNeedId        = util.Error("probe needs an id field")
)
