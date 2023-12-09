package global

import (
	"path/filepath"

	"github.com/metux/go-magicdict/magic"

	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/spec/cache"
	"github.com/metux/go-metabuild/spec/check"
	"github.com/metux/go-metabuild/spec/features"
	"github.com/metux/go-metabuild/spec/generate"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/specobj"
)

type Global struct {
	specobj.SpecObj
}

func (g Global) GetChecks() check.CheckList {
	chklist := g.EntrySpecList(KeyConfigureChecks)
	out := make(check.CheckList, 0, len(chklist))
	for _, ent := range chklist {
		chk := check.Check{
			SpecObj:   ent,
			Cache:     g.GetCache(),
			BuildConf: g.BuildConf()}
		chk.Init()
		out = append(out, chk)
	}
	return out
}

func (g Global) GetGenerates() generate.GenerateList {
	bc := g.BuildConf()
	list := g.EntrySpecList(KeyConfigureGenerate)
	out := make(generate.GenerateList, 0, len(list))
	for _, ent := range list {
		g := generate.Generate{SpecObj: ent, BuildConf: bc}
		g.Init()
		out = append(out, g)
	}
	return out
}

func (g Global) GetTargetObjects() map[string]target.TargetObject {
	bc := g.BuildConf()
	c := g.GetCache()
	m := make(map[string]target.TargetObject)
	// FIXME: use SubSpecList / EntrySpecMap
	for _, n := range g.EntryKeys(KeyTargetObjects) {
		m[string(n)] = target.MakeTargetObject(g.EntrySpec(KeyTargetObjects.Append(n)), n, bc, c)
	}
	return m
}

func (g Global) BuildConf() buildconf.BuildConf {
	return buildconf.BuildConf{SpecObj: g.EntrySpec(KeyBuildConf), Features: g.GetFeatureMap()}
}

func (g Global) GetCache() cache.Cache {
	return cache.NewCache(g.EntrySpec(KeyCache))
}

func (g Global) GetFeatureMap() features.FeatureMap {
	return features.FeatureMap{g.EntrySpec(KeyFeatures)}
}

func (g Global) Init() {
	// init buildconf
	bc := g.BuildConf()
	bc.Init()

	// init checks
	for _, ent := range g.EntrySpecList(KeyConfigureChecks) {
		chk := check.Check{
			SpecObj:   ent,
			Cache:     g.GetCache(),
			BuildConf: g.BuildConf()}
		chk.Init()
	}

	// init generates
	for _, ent := range g.EntrySpecList(KeyConfigureGenerate) {
		g := generate.Generate{SpecObj: ent, BuildConf: bc}
		g.Init()
	}

	// init features
	fm := g.GetFeatureMap()
	fm.Init()
}

func (g Global) PostConfig() {
	for _, ent := range g.GetTargetObjects() {
		ent.LoadTargetDefaults()
	}
}

func LoadGlobal(fn string, dflt string) (Global, error) {
	md, err := magic.YamlLoad(fn, dflt)
	g := Global{SpecObj: specobj.NewSpecObj(md)}

	absfn, _ := filepath.Abs(fn)
	g.DefaultPutStr(KeySysConfigPath, fn)
	g.DefaultPutStr(KeySysConfigDir, filepath.Dir(fn))
	g.DefaultPutStr(KeySysConfigAbsDir, filepath.Dir(absfn))
	g.DefaultPutStr(KeySysConfigAbsPath, absfn)
	g.DefaultPutStr(KeySysConfigBase, filepath.Base(fn))

	absdflt, _ := filepath.Abs(dflt)
	g.DefaultPutStr(KeySysSettingsPath, dflt)
	g.DefaultPutStr(KeySysSettingsDir, filepath.Dir(dflt))
	g.DefaultPutStr(KeySysSettingsAbsDir, filepath.Dir(absdflt))
	g.DefaultPutStr(KeySysSettingsAbsPath, absdflt)
	g.DefaultPutStr(KeySysSettingsBase, filepath.Base(dflt))

	g.Init()

	return g, err
}
