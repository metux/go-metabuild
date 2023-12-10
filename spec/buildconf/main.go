package buildconf

import (
	"strings"

	"github.com/metux/go-metabuild/spec/features"
	"github.com/metux/go-metabuild/util/specobj"
)

type Key = specobj.Key
type SpecObj = specobj.SpecObj

const (
	KeyCheckConfig = Key("config")
	KeyPackages    = Key("packages")
	KeyFlags       = Key("flags")
	KeyVersion     = Key("version")
	KeyAuthors     = Key("authors")
	KeyLinguas     = Key("linguas")
)

type BuildConf struct {
	specobj.SpecObj
	Features features.FeatureMap
}

// define a conditional/switch symbol
func (bc BuildConf) ConfigBool(sym string, val bool) {
	bc.KConfBool(sym, val)
	bc.ACBool(sym, val)
}

func (bc BuildConf) ConfigStr(sym string, val string) {
	bc.KConfStr(sym, val)
	bc.ACStr(sym, val)
}

func (bc BuildConf) ConfigStrList(sym string, val []string) {
	bc.ConfigStr(sym, strings.Join(val, " "))
}

func (bc BuildConf) PkgNameTrans(id string) string {
	return bc.TargetDistro().PackageNameTrans(id)
}

func (bc BuildConf) Init() {
	bc.DefaultPutStr("@features", "${features}")
	bc.DefaultPutStr(KeyBuildDir, BuildDir)
	bc.DefaultPutStr(KeyBuildDirTmp, BuildDirTmp)
	bc.DefaultPutStr(KeyBuildDirDist, BuildDirDist)
}

func (bc BuildConf) Flags(build bool) SpecObj {
	return bc.SubForBuild(build, KeyFlags)
}
