package base

import (
	"fmt"

	"github.com/metux/go-metabuild/util"
	cmdrun "github.com/metux/go-metabuild/util/cmd"
)

const (
	LangC = "C"
)

const (
	CompilerGCC   = "gcc"
	CompilerClang = "clang"
)

const (
	ErrNoUsableCompiler = util.Error("CC no usable compiler")
	ErrCrossMissingCC   = util.Error("when crosscompiling, either HOST_CC must be set or CROSS_COMPILE set to target toolchain prefix")
)

type CompilerInfo struct {
	Language      string
	Id            string
	Command       []string
	Archiver      []string
	Linker        []string
	Version       string
	Machine       Machine
	CrossForHost  bool // this is crosscompiler for foreign target
	CrossForBuild bool // on crosscompile this is the host compiler
	Sysroot       string
	CrossPrefix   string
}

func (ci CompilerInfo) Found() bool {
	return ci.Id != ""
}

func (ci CompilerInfo) String() string {
	return fmt.Sprintf(
		"lang=%s: ID=%s VER=\"%s\" MACH=\"%s\" CMD=%s",
		ci.Language,
		ci.Id,
		ci.Version,
		ci.Machine,
		ci.Command)
}

func (ci *CompilerInfo) RunCatchOut(param string) string {
	out, _ := cmdrun.RunOutOne(append(ci.Command, param), false)
	return out
}
