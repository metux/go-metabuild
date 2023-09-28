package distro

import (
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/specobj"
)

const (
	ErrPkgFormatUnsupported = util.Error("unsupported package format")
)

type Distro struct {
	specobj.SpecObj
	DistroName string
}

func (d Distro) PackageFormat() string {
	return d.EntryStr(KeyPackageFormat)
}

func (d Distro) Name() string {
	return d.DistroName
}

func (d Distro) PackageNameTrans(pkg string) string {
	return d.EntryStr(KeyPackages.AppendStr(pkg).Append("name"))
}

func (d Distro) PackageIds() []Key {
	return d.EntryKeys(KeyPackages)
}

func (d Distro) Packages() []DistPkg {
	return specobj.SpecXformSpecList(
		d.SpecObj,
		KeyPackages,
		func(so specobj.SpecObj) DistPkg { return NewDistPkg(so, d) })
}

func (d Distro) Package(p string) DistPkg {
	return NewDistPkg(d.EntrySpec(KeyPackages.AppendStr(p)), d)
}

func NewDistro(so specobj.SpecObj, name string) Distro {
	return Distro{so, name}
}
