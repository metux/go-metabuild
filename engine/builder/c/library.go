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
}

func (b *BuilderCLibrary) JobPrepare(id jobs.JobId) error {
	cflags := b.CFlags()

	libname := b.RequiredEntryStr(target.KeyLibName)
	pkgname := b.RequiredEntryStr(target.KeyPkgName)
	pkgid := b.RequiredEntryStr(target.KeyLibraryPkgId)

	cflags = []string{"-I."}

	pi := compiler.PkgConfigInfo{
		Name:          pkgname,
		PkgSpec:       pkgname,
		SharedLdflags: []string{"-L.", "-l" + libname},
		SharedCflags:  cflags,
		StaticLdflags: []string{"-L.", "-l:" + b.RequiredEntryStr(target.KeyStaticLib)},
		StaticCflags:  cflags,
	}

	return b.BuildConf.SetPkgConfig(b.ForBuild(), pkgid, pi)
}

func (b BuilderCLibrary) copySub(sub base.BaseBuilder) {
	sub.DefaultPutStrList(target.KeySource, b.FeaturedStrList(target.KeySource))
	sub.DefaultPutStrList(target.KeyCDefines, b.FeaturedStrList(target.KeyCDefines))
	sub.DefaultPutStrList(target.KeyCCflags, b.FeaturedStrList(target.KeyCCflags))
	sub.DefaultPutStrList(target.KeyLinkStatic, b.FeaturedStrList(target.KeyLinkStatic))
	sub.DefaultPutStrList(target.KeyLinkShared, b.FeaturedStrList(target.KeyLinkShared))
	sub.DefaultPutStrList(target.KeyLinkBoth, b.FeaturedStrList(target.KeyLinkBoth))
	sub.DefaultPutStrList(target.KeyPkgconfImport, b.FeaturedStrList(target.KeyPkgconfImport))
	sub.DefaultPutStrList(target.KeyIncludeDirs, b.FeaturedStrList(target.KeyIncludeDirs))
	sub.DefaultPutStrList(target.KeyLibDirs, b.FeaturedStrList(target.KeyLibDirs))
	sub.DefaultPutStrList(target.KeyJobDepends, b.JobDepends())
}

func (b *BuilderCLibrary) mksub1(subkey spec.Key) BaseCBuilder {
	subtarget := b.SubTarget(subkey)
	newbuilder := MakeBaseCBuilder(subtarget, b.JobId()+"/"+string(subkey))
	b.copySub(newbuilder.BaseBuilder)
	return newbuilder
}

func (b *BuilderCLibrary) mkHdrSub(subkey spec.Key) BuilderCLibraryHeaders {
	newbuilder := MakeBaseCBuilder(b.SubTarget(subkey), b.JobId()+"/"+string(subkey))
	newbuilder.SetType(target.TypeCHeader)
	// needs to be explicitly initialized, since not yet known in post-configure phase
	newbuilder.LoadTargetDefaults()
	b.copySub(newbuilder.BaseBuilder)
	return BuilderCLibraryHeaders{newbuilder}
}

// FIXME: support skipping some of them
func (b BuilderCLibrary) JobSub() ([]jobs.Job, error) {

	jobs := []jobs.Job{}

	if !b.EntryBoolDef("skip/shared", false) {
		jobs = append(jobs,
			&BuilderCLibraryShared{b.mksub1("shared")},
			&BuilderCLibraryDevlink{b.mksub1("devlink")})
	}
	if !b.EntryBoolDef("skip/static", false) {
		jobs = append(jobs, &BuilderCLibraryStatic{b.mksub1("static")})
	}
	if !b.EntryBoolDef("skip/pkgconf", false) {
		jobs = append(jobs, &BuilderCLibraryPkgConfig{b.mksub1("pkgconf")})
	}

	for _, h := range b.EntryKeys(target.KeyHeaders) {
		jobs = append(jobs, b.mkHdrSub(target.KeyHeaders.Append(h)))
	}
	return jobs, nil
}

func MakeCLibrary(o spec.TargetObject, id string) *BuilderCLibrary {
	b := BuilderCLibrary{BaseCBuilder: MakeBaseCBuilder(o, id)}
	return &b
}
