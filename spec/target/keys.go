package target

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	KeyId   = Key("id")
	KeyType = Key("type")

	// build host vs target host
	KeyBuild = Key("build")

	KeyCompilerLang = Key("compiler/lang")

	// all targets: installation
	KeyInstall        = Key("install")
	KeyInstallDir     = Key("install/dir")
	KeyInstallSubdir  = Key("install/subdir")
	KeyInstallPackage = Key("install/package")
	KeyInstallPerm    = Key("install/perm")

	// library main
	KeyHeaders   = Key("headers")
	KeyMapfile   = Key("mapfile")
	KeyLibName   = Key("libname")
	KeyStaticLib = Key("static/libfile")

	// library devlink
	KeyLinkTarget = Key("target")

	KeySource       = Key("source")
	KeySourceDir    = Key("source/dir")
	KeySourceSuffix = Key("source/suffix")

	KeySubPkg  = Key("subpkg")
	KeyName    = Key("name")
	KeyFile    = Key("file")
	KeySymlink = Key("symlink")

	KeyIncludeDirs = Key("include-dirs")
	KeyLibDirs     = Key("lib-dirs")

	KeyLinkStatic = Key("link/static")
	KeyLinkShared = Key("link/shared")
	KeyLinkBoth   = Key("link/both")

	KeyImport = Key("import")
)

// locales
const (
	KeyI18nLinguas  = Key("i18n/linguas")
	KeyI18nCategory = Key("i18n/category")
	KeyI18nDomain   = Key("i18n/domain")
	KeyI18nPoDir    = Key("i18n/po/dir")
)

// generators
const (
	KeyResourceDir  = Key("resource/dir")
	KeyResourceName = Key("resource/name")

	KeyOutputCHeader   = Key("output/c/header")
	KeyOutputCSource   = Key("output/c/source")
	KeyOutputGResource = Key("output/gresource")
	KeyOutputSuffix    = Key("output/suffix")
)

// C specific settings
const (
	KeyCDefines = Key("c/defines")
	KeyCCflags  = Key("c/cflags")
	KeyCLdflags = Key("c/ldflags")
)

// desktop file settings
const (
	KeyDesktopType        = Key("desktop/type")
	KeyDesktopName        = Key("desktop/name")
	KeyDesktopGenericName = Key("desktop/genericname")
	KeyDesktopComment     = Key("desktop/comment")
	KeyDesktopIconFile    = Key("desktop/icon-file")
	KeyDesktopExec        = Key("desktop/exec")
	KeyDesktopTryExec     = Key("desktop/tryexec")
	KeyDesktopTerminal    = Key("desktop/terminal")
	KeyDesktopCategories  = Key("desktop/categories")
)
