package target

import (
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/compiler"
)

// pkgconfig settings
const (
	// importing by pkgconf from within a target
	KeyPkgconfImport = Key("pkgconf-import")

	// pkgconf descriptor attributes
	KeyPkgName         = Key("name")
	KeyPkgVersion      = Key("version")
	KeyPkgDescription  = Key("description")
	KeyPkgPrefix       = Key("prefix")
	KeyPkgExecPrefix   = Key("exec-prefix")
	KeyPkgLibdir       = Key("libdir")
	KeyPkgIncludedir   = Key("includedir")
	KeyPkgSharedLibdir = Key("sharedlibdir")
	KeyPkgArchive      = Key("archive")
	KeyPkgLibname      = Key("libname")
	KeyPkgPackage      = Key("package")
)

func (o TargetObject) tryImport(id string) compiler.PkgConfigInfo {
	forBuild := o.ForBuild()
	pkg := o.BuildConf.PkgConfig(forBuild, id)
	if pkg.Valid() {
		return pkg
	}
	return pkg
}

func (o TargetObject) PkgImports() []compiler.PkgConfigInfo {
	imports := o.FeaturedStrList(KeyPkgconfImport)
	pkgs := make([]compiler.PkgConfigInfo, len(imports))
	for idx, i := range imports {
		pkgs[idx] = o.tryImport(i)
		if !pkgs[idx].Valid() {
			util.Panicf("config error: missing pkg-config check for: %s", pkgs[idx].Id)
		}
	}
	return pkgs
}
