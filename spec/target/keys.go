package target

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	// internal keys -- automatically created after loading
	KeyInternId       = Key("@id")
	KeyInternIdSuffix = Key("@id/suffix")
	KeyInternType     = Key("@type")
	KeyInternBasename = Key("@basename")
	KeyInternDirname  = Key("@dirname")

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

	KeySource       = Key("source")
	KeySourceDir    = Key("source/dir")
	KeySourceSuffix = Key("source/suffix")

	KeyName       = Key("name")
	KeyFile       = Key("file")
	KeyJobDepends = Key("job/depends")

	KeyIncludeDir = Key("include/dir")
	KeyLibDirs    = Key("lib-dirs")

	KeyLinkStatic = Key("link/static")
	KeyLinkShared = Key("link/shared")
	KeyLinkBoth   = Key("link/both")
)

// libraries (main)
const (
	// pkgconf ID for internal libs
	KeyLibraryPkgId   = Key("library/pkgid")
	KeyLibraryMapFile = Key("library/mapfile")
	KeyHeaders        = Key("headers")
	KeyLibraryName    = Key("library/name")
	KeyStaticLib      = Key("static::file")
	KeyLibraryDir     = Key("library/dir")

	// static / archives
	KeyLibraryLinkWhole = Key("library/link-whole")

	// library devlink
	KeyLinkTarget = Key("target")
)

// internal builder configuration
const (
	KeyBuilderCommand = Key("builder/command")
	KeyBuilderDriver  = Key("builder/driver")
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
	KeyOutputFormat    = Key("output/format")
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

// manpages
const (
	KeyManpageSection  = Key("man/section")
	KeyManpageAlias    = Key("man/alias")
	KeyManpageCompress = Key("man/compress")
)
