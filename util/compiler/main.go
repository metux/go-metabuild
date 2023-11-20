package compiler

import (
	"fmt"

	"github.com/metux/go-metabuild/util/compiler/base"
	"github.com/metux/go-metabuild/util/compiler/gnu"
	"github.com/metux/go-metabuild/util/compiler/probe"
)

const (
	LangC   = base.LangC
	LangCxx = base.LangCxx
)

type (
	BinaryFileInfo = base.BinaryFileInfo
	CompilerArg    = base.CompilerArg
	CompilerInfo   = base.CompilerInfo
	PkgConfigInfo  = base.PkgConfigInfo
	CCompiler      = base.CCompiler
)

var (
	ParseMachine   = base.ParseMachine
	DetectCC       = probe.DetectCC
	DetectCXX      = probe.DetectCXX
	PkgConfigQuery = base.PkgConfigQuery
)

func NewCCompiler(ci base.CompilerInfo, tempdir string) base.CCompiler {
	switch ci.Id {
	// assume gcc and clang use the same cmdline args
	case base.CompilerGCC, base.CompilerClang:
		return gnu.NewCCompiler(ci, tempdir)
	}
	panic(fmt.Sprintf("unsupported compiler type: %s", ci.Id))
	return nil
}
