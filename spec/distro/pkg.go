package distro

import (
	"log"

	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/pkgmgr"
	"github.com/metux/go-metabuild/util/specobj"
)

const (
	KeyPkgName        = Key("name")
	KeyPkgVersion     = Key("version")
	KeyPkgDescription = Key("description")
	KeyPkgArch        = Key("arch")
	KeyPkgMaintainer  = Key("maintainer")
	KeyPkgSection     = Key("section")
	KeyPkgLocalDepend = Key("local-depend")
	KeyPkgExtdepend   = Key("pkg-depend")
	KeyPkgBugs        = Key("bugs")
	KeyPkgHomepage    = Key("homepage")
	KeyPkgOrigin      = Key("origin")
	KeyPkgPriority    = Key("priority")
	KeyPkgSkip        = Key("skip")
)

type DistPkg struct {
	specobj.SpecObj
	Distro Distro
}

func (pkg DistPkg) Id() string {
	return pkg.EntryStr(Key("@@KEY"))
}

func (pkg DistPkg) DistroName() string {
	return pkg.Distro.Name()
}

func (pkg DistPkg) Name() string {
	return pkg.EntryStr(KeyPkgName)
}

func (pkg DistPkg) Version() string {
	return pkg.EntryStr(KeyPkgVersion)
}

func (pkg DistPkg) Description() string {
	return pkg.EntryStr(KeyPkgDescription)
}

func (pkg DistPkg) Arch() string {
	return pkg.EntryStr(KeyPkgArch)
}

func (pkg DistPkg) Maintainer() string {
	return pkg.EntryStr(KeyPkgMaintainer)
}

func (pkg DistPkg) Section() string {
	return pkg.EntryStr(KeyPkgSection)
}

func (pkg DistPkg) Depends() []string {
	return pkg.EntryStrList(KeyPkgLocalDepend)
}

func (pkg DistPkg) ControlInfo(deps []string) pkgmgr.PkgControl {

	// add our local deps
	for _, d := range pkg.Depends() {
		p := pkg.Distro.Package(d)
		if p.Skipped() {
			log.Println("skipping package:", p.Name())
		} else {
			deps = append(deps, p.Name()+" (>="+p.Version()+")")
		}
	}

	return pkgmgr.PkgControl{
		Package:      pkg.EntryStr(KeyPkgName),
		Version:      pkg.EntryStr(KeyPkgVersion),
		Description:  pkg.EntryStr(KeyPkgDescription),
		Architecture: pkg.EntryStr(KeyPkgArch),
		Maintainer:   pkg.EntryStr(KeyPkgMaintainer),
		Section:      pkg.EntryStr(KeyPkgSection),
		Origin:       pkg.EntryStr(KeyPkgOrigin),
		Homepage:     pkg.EntryStr(KeyPkgHomepage),
		Bugs:         pkg.EntryStr(KeyPkgBugs),
		Priority:     pkg.EntryStr(KeyPkgPriority),
		Depend:       util.Uniq(pkg.EntryStrList("pkg-depend"), deps),
	}
}

func (pkg DistPkg) Skipped() bool {
	return pkg.EntryBoolDef(KeyPkgSkip, false)
}

func NewDistPkg(so specobj.SpecObj, d Distro) DistPkg {
	return DistPkg{so, d}
}
