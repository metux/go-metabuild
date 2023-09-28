package buildconf

import (
	"github.com/metux/go-metabuild/util/compiler"
	"github.com/metux/go-metabuild/util/specobj"
)

const (
	KeyForLang = Key("lang")
)

// Suffixes for per-language compiler attributes
const (
	KeyCompilerTarget        = Key("target")
	KeyCompilerArch          = Key("arch")
	KeyCompilerVendor        = Key("vendor")
	KeyCompilerKernel        = Key("kernel")
	KeyCompilerSystem        = Key("system")
	KeyCompilerCommand       = Key("command")
	KeyCompilerArchiver      = Key("archiver")
	KeyCompilerId            = Key("id")
	KeyCompilerCrossForHost  = Key("cross-for-host")
	KeyCompilerCrossForBuild = Key("cross-for-build")
	KeyCompilerCrossPrefix   = Key("cross-prefix")
	KeyCompilerSysroot       = Key("sysroot")
)

func (bc BuildConf) CompilerSub(build bool, lang string) specobj.SpecObj {
	return bc.SubForBuild(build, KeyForLang.AppendStr(lang))
}

// Store CompilerInfo in BuildConfig
// build - true if it's the compiler for the build system (different from host on crosscompile)
func (bc BuildConf) SetCompilerInfo(build bool, ci compiler.CompilerInfo) {
	sub := bc.CompilerSub(build, ci.Language)
	sub.EntryPutBool(KeyCompilerCrossForHost, ci.CrossForHost)
	sub.EntryPutBool(KeyCompilerCrossForBuild, ci.CrossForBuild)
	sub.EntryPutStr(KeyCompilerCrossPrefix, ci.CrossPrefix)
	sub.EntryPutStrList(KeyCompilerArchiver, ci.Archiver)
	sub.EntryPutStr(KeyCompilerTarget, ci.Machine.String())
	sub.EntryPutStr(KeyCompilerArch, ci.Machine.Arch)
	sub.EntryPutStr(KeyCompilerVendor, ci.Machine.Vendor)
	sub.EntryPutStr(KeyCompilerKernel, ci.Machine.Kernel)
	sub.EntryPutStr(KeyCompilerSystem, ci.Machine.System)
	sub.EntryPutStr(KeyCompilerId, ci.Id)
	sub.EntryPutStr(KeyCompilerSysroot, ci.Sysroot)
	sub.EntryPutStrList(KeyCompilerCommand, ci.Command)
}

func (bc BuildConf) CompilerInfo(build bool, lang string) compiler.CompilerInfo {
	sub := bc.CompilerSub(build, lang)
	return compiler.CompilerInfo{
		Language:      lang,
		CrossForHost:  sub.EntryBoolDef(KeyCompilerCrossForHost, false),
		CrossForBuild: sub.EntryBoolDef(KeyCompilerCrossForBuild, false),
		CrossPrefix:   sub.EntryStr(KeyCompilerCrossPrefix),
		Machine:       compiler.ParseMachine(sub.EntryStr(KeyCompilerTarget)),
		Command:       sub.EntryStrList(KeyCompilerCommand),
		Archiver:      sub.EntryStrList(KeyCompilerArchiver),
		Id:            sub.EntryStr(KeyCompilerId),
		Sysroot:       sub.EntryStr(KeyCompilerSysroot),
	}
}
