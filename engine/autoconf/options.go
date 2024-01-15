package autoconf

import (
	"os"
	"strings"

	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/global"
	"github.com/metux/go-metabuild/spec/options"
	"github.com/metux/go-metabuild/util"
)

type Key = spec.Key

// FIXME: move this out to frontend code
// FIXME: move env lookup into OptionMap, but pass in env []
func optionsFromEnv(cf global.Global, fm options.OptionMap) {
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

func optionsProcess(cf global.Global, fm options.OptionMap) error {
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
					return util.ConfigError("missing package import \"%s\" for option \"%s\"", pkg, f.Id)
				}
			}
		}
	}
	return nil
}

func RunConfigureOptions(cf global.Global) error {
	fmap := cf.GetOptionMap()
	optionsFromEnv(cf, fmap)
	return optionsProcess(cf, fmap)
}
