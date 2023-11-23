package autoconf

import (
	"os"
	"strings"

	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/features"
	"github.com/metux/go-metabuild/spec/global"
	"github.com/metux/go-metabuild/util"
)

type Key = spec.Key

// FIXME: move this out to frontend code
// FIXME: move env lookup into FeatureMap, but pass in env []
func featuresFromEnv(cf global.Global, fm features.FeatureMap) {
	for _, f := range fm.All() {
		fup := strings.ToUpper(string(f.Id))
		if v := os.Getenv("ENABLE_" + fup); v != "" {
			f.SetOn()
		}
		if v := os.Getenv("DISABLE_" + fup); v != "" {
			f.SetOff()
		}
		if v := os.Getenv("WITH_" + fup); v != "" {
			f.Set(v)
		}
	}
}

func featuresProcess(cf global.Global, fm features.FeatureMap) {
	bc := cf.BuildConf()

	subBuild := bc.SubForBuild(true, "flags")
	subHost := bc.SubForBuild(false, "flags")

	for _, f := range fm.All() {
		flags := f.FlagsOn()
		for _, k := range flags.Keys() {
			data := flags.EntryStrList(k)
			subBuild.EntryStrListAppendList(k, data)
			subHost.EntryStrListAppendList(k, data)
		}
		if f.ValueYN() == "y" {
			req := f.PkgconfRequire()
			for _, pkg := range req {
				if !bc.PkgConfig(true, pkg).Valid() && !bc.PkgConfig(false, pkg).Valid() {
					util.Panicf("config error: missing package %s for feature %s", pkg, f.Id)
				}
			}
		}
	}
}

func RunConfigureFeatures(cf global.Global) error {
	fmap := cf.GetFeatureMap()
	featuresFromEnv(cf, fmap)
	featuresProcess(cf, fmap)
	return nil
}
