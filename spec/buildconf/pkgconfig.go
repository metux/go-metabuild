package buildconf

import (
	"fmt"
	"strings"

	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/compiler"
)

// suffixes for pkg-config
const (
	KeyPkgSpec          = Key("pkgspec")
	KeyPkgName          = Key("pkg")
	KeyPkgVersion       = Key("version")
	KeyPkgCflagsShared  = Key("shared/cflags")
	KeyPkgCflagsStatic  = Key("static/cflags")
	KeyPkgLdflagsShared = Key("shared/ldflags")
	KeyPkgLdflagsStatic = Key("static/ldflags")
)

// subtree of pkgconfig subtree underneath the build/host subtree
const (
	KeyPkgSub = Key("pkg")
)

func (bc BuildConf) PkgConfigSub(build bool, pkg string) SpecObj {
	return bc.SubForBuild(build, KeyPkgSub.AppendStr(pkg))
}

func (bc BuildConf) SetPkgConfig(build bool, id string, pi compiler.PkgConfigInfo) error {
	sub := bc.PkgConfigSub(build, id)
	sub.EntryPutStr(KeyPkgSpec, pi.PkgSpec)
	sub.EntryPutStr(KeyPkgName, pi.Name)
	sub.EntryPutStr(KeyPkgVersion, pi.Version)
	sub.EntryPutStrList(KeyPkgCflagsShared, pi.SharedCflags)
	sub.EntryPutStrList(KeyPkgCflagsStatic, pi.StaticCflags)
	sub.EntryPutStrList(KeyPkgLdflagsShared, pi.SharedLdflags)
	sub.EntryPutStrList(KeyPkgLdflagsStatic, pi.StaticLdflags)
	sub.EntryPutStrMap("variables", pi.Variables)

	// write out to config
	bc.ConfigBool("HAVE_PKG_"+id, pi.Name != "")
	pkg_prefix := fmt.Sprintf("PKG_%s_%s_", util.ValIf(build, "BUILD", "HOST"), strings.ToUpper(id))
	bc.ConfigStrList(pkg_prefix+"STATIC_CFLAGS", pi.StaticCflags)
	bc.ConfigStrList(pkg_prefix+"STATIC_LDFLAGS", pi.StaticLdflags)
	bc.ConfigStr(pkg_prefix+"VERSION", pi.Version)
	bc.ConfigStrList(pkg_prefix+"SHARED_CFLAGS", pi.SharedCflags)
	bc.ConfigStrList(pkg_prefix+"SHARED_LDFLAGS", pi.SharedLdflags)

	return nil
}

func pkgconf(sub SpecObj, id string) compiler.PkgConfigInfo {
	return compiler.PkgConfigInfo{
		Id:            id,
		PkgSpec:       sub.EntryStr(KeyPkgSpec),
		Name:          sub.EntryStr(KeyPkgName),
		Version:       sub.EntryStr(KeyPkgVersion),
		SharedCflags:  sub.EntryStrList(KeyPkgCflagsShared),
		StaticCflags:  sub.EntryStrList(KeyPkgCflagsStatic),
		SharedLdflags: sub.EntryStrList(KeyPkgLdflagsShared),
		StaticLdflags: sub.EntryStrList(KeyPkgLdflagsStatic),
	}
}

func (bc BuildConf) PkgConfig(build bool, id string) compiler.PkgConfigInfo {
	pkg := pkgconf(bc.PkgConfigSub(build, id), id)
	if !pkg.Valid() {
		// try to get it from target platform
		if !build {
			pkg = pkgconf(bc.EntrySpec(KeyTargetPlatform.Append(KeyPkgSub).AppendStr(id)), id)
		} // FIXME: also should have @host-platform ?
	}
	return pkg
}
