package c

import (
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
	b.LoadTargetDefaults(spec.Key(b.Type()))

	cdefs := b.CDefines()
	ci := b.BuildConf.CompilerInfo(b.ForBuild(), b.CompilerLang())
	jobdep := b.JobDepends()

	// we NEED to initialize them all, but only add them if not skipped
	b.subShared = &BuilderCLibraryShared{b.mksub("shared", target.TypeCLibraryShared, ci, cdefs, jobdep)}
	b.subStatic = &BuilderCLibraryStatic{b.mksub("static", target.TypeCLibraryStatic, ci, cdefs, jobdep)}
	b.subDevlink = &BuilderCLibraryDevlink{b.mksub("devlink", target.TypeCLibraryDevlink, ci, cdefs, jobdep)}
	b.subPkgconfig = &BuilderCLibraryPkgConfig{b.mksub("pkgconf", target.TypeCLibraryPkgconf, ci, cdefs, jobdep)}

	libname := b.RequiredEntryStr(target.KeyLibName)
	pkgname := b.RequiredEntryStr(target.KeyPkgName)
	cflags := []string{"-I."}

	pi := compiler.PkgConfigInfo{
		Name:          pkgname,
		PkgSpec:       pkgname,
		SharedLdflags: []string{"-L.", "-l" + libname},
		SharedCflags:  cflags,
		StaticLdflags: []string{"-L.", "-l:" + b.RequiredEntryStr(target.KeyStaticLib)},
		StaticCflags:  cflags,
	}

	return b.BuildConf.SetPkgConfig(b.ForBuild(), string(b.MyKey()), pi)
}

func (b *BuilderCLibrary) mksub(sub spec.Key, typ spec.Key, ci compiler.CompilerInfo, cdefs []string, jobdep []jobs.JobId) CommonCBuilder {
	newbuilder := base.NewBaseBuilder(b.SubTarget(sub), b.JobId()+"/"+string(sub))
	newbuilder.LoadTargetDefaults(typ)
	return CommonCBuilder{newbuilder, ci, cdefs, jobdep, &b.BaseCBuilder}
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

	cdefs := b.CDefines()
	ci := b.BuildConf.CompilerInfo(b.ForBuild(), b.CompilerLang())
	jobdep := b.JobDepends()
	for _, h := range b.EntryKeys("headers") {
		jobs = append(jobs, BuilderCLibraryHeaders{b.mksub(target.KeyHeaders.Append(h), target.TypeCHeader, ci, cdefs, jobdep)})
	}
	return jobs, nil
}

func MakeBuilderCLibrary(o spec.TargetObject, id string) *BuilderCLibrary {
	b := BuilderCLibrary{BaseCBuilder: NewBaseCBuilder(o, id)}
	b.JobPrepare(id)
	return &b
}
