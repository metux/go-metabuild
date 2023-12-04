package target

import (
	"os"
	"path/filepath"

	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/spec/cache"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/fileutil"
	"github.com/metux/go-metabuild/util/specobj"
)

type TargetObject struct {
	specobj.SpecObj
	BuildConf buildconf.BuildConf
	Cache     cache.Cache
}

func (o TargetObject) FeaturedStrList(k Key) []string {
	ret := o.EntryStrList(k)
	for _, f := range o.BuildConf.Features.All() {
		k2 := Key(string(k) + "@feature/" + string(f.Id) + "=" + f.ValueYN())
		ret = append(ret, o.EntryStrList(k2)...)
	}
	return ret
}

func (o TargetObject) Sources() []string {
	res := []string{}
	for _, f := range util.StrDirPrefix(o.EntryStr(KeySourceDir), o.FeaturedStrList(KeySource)) {
		files, err := filepath.Glob(f)
		util.ErrPanicf(err, "file glob error")
		if len(files) == 0 {
			panic("broken source glob: " + f)
		}
		res = append(res, files...)
	}
	return res
}

func (o TargetObject) RequiredSources() []string {
	s := util.StrDirPrefix(o.EntryStr(KeySourceDir), o.FeaturedStrList(KeySource))
	if len(s) == 0 {
		panic("source attribute required")
	}
	return s
}

func (o TargetObject) RequiredSourceAbs() string {
	s := o.RequiredSources()
	if len(s) != 1 {
		panic("only exactly one source file required")
	}
	a, err := filepath.Abs(s[0])
	util.ErrPanicf(err, "failed retrieving absolute path of %s", s[0])
	return a
}

func (o TargetObject) ForBuild() bool {
	return o.EntryBoolDef(Key("build"), false)
}

func (o TargetObject) MyId() string {
	return o.EntryStr(KeyInternId)
}

func (o TargetObject) LoadTargetDefaults() {
	id := string(o.MyKey())
	o.DefaultPutStr(KeyInternId, id)

	if ext := filepath.Ext(id); ext != "" {
		o.DefaultPutStr(KeyInternIdSuffix, ext[1:])
	}

	k := buildconf.KeyTargetPlatform.Append("targets").Append(o.Type()).MagicLiteralPost()
	o.DefaultPutStrMap(o.BuildConf.EntryStrMap(k))
}

func (o TargetObject) CDefines() []string {
	return o.FeaturedStrList(KeyCDefines)
}

func (o TargetObject) CFlags() []string {
	return o.FeaturedStrList(KeyCCflags)
}

func (o TargetObject) GetFMode(k Key) os.FileMode {
	n, _ := fileutil.FileModeParse(o.EntryStr(k))
	return n
}

func (o TargetObject) InstallPerm() os.FileMode {
	m, err := fileutil.FileModeParse(o.RequiredEntryStr(KeyInstallPerm))
	if m == 0 || err != nil {
		util.Panicf("config error: target's mode=0 or error %s", err)
	}
	return m
}

func (o TargetObject) InstallDir() string {
	return o.RequiredEntryStr(KeyInstallDir) + "/" + o.EntryStr(KeyInstallSubdir)
}

func (o TargetObject) PackageId(scope string) string {
	k := Key(scope + string(KeyInstallPackage))
	id := o.EntryStr(k)
	if id == "" {
		panic("config error: PackageId empty")
	}
	return id
}

func (o TargetObject) CompilerLang() string {
	return o.RequiredEntryStr(KeyCompilerLang)
}

func (o TargetObject) Type() Key {
	return Key(o.EntryStr(KeyType))
}

func (o TargetObject) SetType(k Key) {
	o.EntryPutStr(KeyType, string(k))
}

func (o TargetObject) SubTarget(k specobj.Key) TargetObject {
	return NewTargetObject(o.EntrySpec(k), k, o.BuildConf, o.Cache)
}

func (o TargetObject) WantInstall() bool {
	return o.EntryBoolDef(KeyInstall, true)
}

func (b TargetObject) EntryPathList(k Key) []string {
	list := b.EntryStrList(k)
	for idx, ent := range list {
		d, err := filepath.Abs(ent)
		util.ErrPanicf(err, "error retrieving abspath for %s", d)
		list[idx] = d
	}
	return list
}

func NewTargetObject(spec specobj.SpecObj, k Key, bc buildconf.BuildConf, c cache.Cache) TargetObject {
	obj := TargetObject{
		SpecObj:   spec,
		Cache:     c,
		BuildConf: bc,
	}
	obj.EntryPutStr(KeyInternId, string(k))
	return obj
}
