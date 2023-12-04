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
	cdefs := b.CDefines()
	cflags := b.CFlags()
	ci := b.BuildConf.CompilerInfo(b.ForBuild(), b.CompilerLang())
	jobdep := b.JobDepends()

	tShared := target.TypeCLibraryShared
	tStatic := target.TypeCLibraryStatic
	tDevlink := target.TypeCLibraryDevlink
	tPkgconf := target.TypeCLibraryPkgconf

	switch lang := b.CompilerLang(); lang {
	case compiler.LangC:
	case compiler.LangCxx:
		tShared = target.TypeCxxLibraryShared
		tStatic = target.TypeCxxLibraryStatic
		tDevlink = target.TypeCxxLibraryDevlink
		tPkgconf = target.TypeCxxLibraryPkgconf
	default:
		panic(fmt.Sprintf("unsupported lang: %s", lang))
	}

	// we NEED to initialize them all, but only add them if not skipped
	b.subShared = &BuilderCLibraryShared{b.mksub("shared", tShared, ci, cdefs, cflags, jobdep)}
	b.subStatic = &BuilderCLibraryStatic{b.mksub("static", tStatic, ci, cdefs, cflags, jobdep)}
	b.subDevlink = &BuilderCLibraryDevlink{b.mksub("devlink", tDevlink, ci, cdefs, cflags, jobdep)}
	b.subPkgconfig = &BuilderCLibraryPkgConfig{b.mksub("pkgconf", tPkgconf, ci, cdefs, cflags, jobdep)}

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

func (b *BuilderCLibrary) mksub(sub spec.Key, typ spec.Key, ci compiler.CompilerInfo, cdefs []string, cflags []string, jobdep []jobs.JobId) CommonCBuilder {
	newbuilder := base.NewBaseBuilder(b.SubTarget(sub), b.JobId()+"/"+string(sub))
	newbuilder.SetType(typ)
	// needs to be explicitly initialized, since not yet known in post-configure phase
	newbuilder.LoadTargetDefaults()
	return CommonCBuilder{newbuilder, ci, cdefs, cflags, jobdep, &b.BaseCBuilder}
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

	cdefs := b.CDefines()
	cflags := b.CFlags()
	ci := b.BuildConf.CompilerInfo(b.ForBuild(), b.CompilerLang())
	jobdep := b.JobDepends()
	for _, h := range b.EntryKeys(target.KeyHeaders) {
		jobs = append(jobs, BuilderCLibraryHeaders{b.mksub(target.KeyHeaders.Append(h), t, ci, cdefs, cflags, jobdep)})
	}
	return jobs, nil
}

func MakeBuilderCLibrary(o spec.TargetObject, id string) *BuilderCLibrary {
	b := BuilderCLibrary{BaseCBuilder: NewBaseCBuilder(o, id)}
	b.JobPrepare(id)
	return &b
}
