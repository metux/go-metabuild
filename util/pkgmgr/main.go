package pkgmgr

import (
	"github.com/metux/go-metabuild/util/pkgmgr/base"
	"github.com/metux/go-metabuild/util/pkgmgr/dpkg"
)

type PkgFileEntry = base.PkgFileEntry

type PkgControl = base.PkgControl

type PackageManager = base.PackageManager

var (
	NewDpkg         = dpkg.NewDpkg
	NewPkgFileEntry = base.NewPkgFileEntry
)
