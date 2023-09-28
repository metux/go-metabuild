package autoconf

import (
	"log"
	"os"
	"strings"

	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/features"
	"github.com/metux/go-metabuild/spec/global"
)

type Key = spec.Key

// FIXME: move this out to frontend code
// FIXME: move env lookup into FeatureMap, but pass in env []
func featuresFromEnv(cf global.Global, fm features.FeatureMap) {
	for _, f := range fm.All() {
		fup := strings.ToUpper(string(f.Id))
		log.Println("checking feature", f.Id, fup)
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
	}
}

func RunConfigureFeatures(cf global.Global) error {
	fmap := cf.GetFeatureMap()
	featuresFromEnv(cf, fmap)
	featuresProcess(cf, fmap)
	return nil
}
