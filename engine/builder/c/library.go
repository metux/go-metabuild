package c

import (
	"fmt"

	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/compiler"
	"github.com/metux/go-metabuild/util/jobs"
)

type BuilderCLibrary struct {
	BaseCBuilder

	subShared    *BuilderCLibraryShared
	subStatic    *BuilderCLibraryStatic
	subDevlink   *BuilderCLibraryDevlink
	subPkgconfig *BuilderCLibraryPkgConfig
}

func (b *BuilderCLibrary) JobPrepare(id jobs.JobId) error {
	cflags := b.CFlags()

	// we NEED to initialize them all, but only add them if not skipped
	b.subShared = &BuilderCLibraryShared{b.mksub1("shared")}
	b.subStatic = &BuilderCLibraryStatic{b.mksub1("static")}
	b.subDevlink = &BuilderCLibraryDevlink{b.mksub1("devlink")}
	b.subPkgconfig = &BuilderCLibraryPkgConfig{b.mksub1("pkgconf")}

	libname := b.RequiredEntryStr(target.KeyLibName)
	pkgname := b.RequiredEntryStr(target.KeyPkgName)
	cflags = []string{"-I."}

	pi := compiler.PkgConfigInfo{
		Name:          pkgname,
		PkgSpec:       pkgname,
		SharedLdflags: []string{"-L.", "-l" + libname},
		SharedCflags:  cflags,
		StaticLdflags: []string{"-L.", "-l:" + b.RequiredEntryStr(target.KeyStaticLib)},
		StaticCflags:  cflags,
	}

	return b.BuildConf.SetPkgConfig(b.ForBuild(), b.MyId(), pi)
}

func (b *BuilderCLibrary) mksub1(sub spec.Key) CommonCBuilder {
	newbuilder := base.MakeBaseBuilder(b.SubTarget(sub), b.JobId()+"/"+string(sub))
	return CommonCBuilder{newbuilder, &b.BaseCBuilder}
}

func (b *BuilderCLibrary) mkHdrSub(sub spec.Key, typ spec.Key) BuilderCLibraryHeaders {
	newbuilder := base.MakeBaseBuilder(b.SubTarget(sub), b.JobId()+"/"+string(sub))
	newbuilder.SetType(typ)
	// needs to be explicitly initialized, since not yet known in post-configure phase
	newbuilder.LoadTargetDefaults()
	return BuilderCLibraryHeaders{CommonCBuilder{newbuilder, &b.BaseCBuilder}}
}

// FIXME: support skipping some of them
func (b BuilderCLibrary) JobSub() ([]jobs.Job, error) {

	jobs := []jobs.Job{}

	if !b.EntryBoolDef("skip/shared", false) {
		jobs = append(jobs, b.subShared, b.subDevlink)
	}
	if !b.EntryBoolDef("skip/static", false) {
		jobs = append(jobs, b.subStatic)
	}
	if !b.EntryBoolDef("skip/pkgconf", false) {
		jobs = append(jobs, b.subPkgconfig)
	}

	t := target.TypeCHeader
	switch lang := b.CompilerLang(); lang {
	case compiler.LangC:
	case compiler.LangCxx:
		t = target.TypeCxxHeader
	default:
		panic(fmt.Sprintf("unsupported lang: %s", lang))
	}

	for _, h := range b.EntryKeys(target.KeyHeaders) {
		jobs = append(jobs, b.mkHdrSub(target.KeyHeaders.Append(h), t))
	}
	return jobs, nil
}

func MakeCLibrary(o spec.TargetObject, id string) *BuilderCLibrary {
	b := BuilderCLibrary{BaseCBuilder: MakeBaseCBuilder(o, id)}
	return &b
}
