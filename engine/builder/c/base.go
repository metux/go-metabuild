package c

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/compiler"
	"github.com/metux/go-metabuild/util/jobs"
)

type BaseCBuilder struct {
	base.BaseBuilder
}

// FIXME: move this to individual builders ?
func (b BaseCBuilder) JobDepends() []jobs.JobId {
	j := b.BaseBuilder.JobDepends()
	for _, x := range b.EntryStrList(target.KeyLinkStatic) {
		j = append(j, jobs.JobId(x))
	}
	for _, x := range b.EntryStrList(target.KeyLinkShared) {
		j = append(j, jobs.JobId(x))
	}
	for _, x := range b.EntryStrList(target.KeyLinkBoth) {
		j = append(j, jobs.JobId(x))
	}
	return j
}

func (b BaseCBuilder) ImportInternalLib(id string, wantShared bool, wantStatic bool) compiler.PkgConfigInfo {
	pi := b.BuildConf.PkgConfig(b.ForBuild(), id)
	pi.WantShared = wantShared
	pi.WantStatic = wantStatic
	return pi
}

func (b BaseCBuilder) ImportSrcdir() compiler.PkgConfigInfo {
	cflags := append(b.FeaturedStrList(target.KeyCCflags), util.StrPrefix("-I", b.EntryStrList(target.KeyIncludeDirs))...)
	ldflags := append(b.FeaturedStrList(target.KeyCLdflags), util.StrPrefix("-L", b.EntryStrList(target.KeyLibDirs))...)

	return compiler.PkgConfigInfo{
		Private:       true,
		SharedCflags:  cflags,
		SharedLdflags: ldflags,
		StaticCflags:  cflags,
		StaticLdflags: ldflags,
	}
}

func (b BaseCBuilder) AllImports() []compiler.PkgConfigInfo {
	imports := []compiler.PkgConfigInfo{}

	for _, x := range b.EntryStrList(target.KeyLinkStatic) {
		imports = append(imports, b.ImportInternalLib(x, false, true))
	}
	for _, x := range b.EntryStrList(target.KeyLinkShared) {
		imports = append(imports, b.ImportInternalLib(x, true, false))
	}
	for _, x := range b.EntryStrList(target.KeyLinkBoth) {
		imports = append(imports, b.ImportInternalLib(x, false, false))
	}

	imports = append(imports, b.ImportSrcdir())

	pkgs := b.PkgImports()
	// sanity check for pkg-config imports
	for _, p := range pkgs {
		if !p.Valid() {
			util.Panicf("config error: missing pkg-config check for: %s", p.Id)
		}
	}

	imports = append(imports, pkgs...)
	return imports
}

func MakeBaseCBuilder(o spec.TargetObject, id string) BaseCBuilder {
	return BaseCBuilder{base.NewBaseBuilder(o, id)}
}
