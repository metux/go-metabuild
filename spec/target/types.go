package target

// target types
const (
	/* plain C */
	TypeCLibrary        = Key("c/library")
	TypeCExecutable     = Key("c/executable")
	TypeCHeader         = Key("c/header")
	TypeCLibraryStatic  = Key("c/library/static")
	TypeCLibraryShared  = Key("c/library/shared")
	TypeCLibraryPkgconf = Key("c/library/pkgconf")
	TypeCLibraryDevlink = Key("c/library/devlink")

	/* C++ */
	TypeCxxExecutable     = Key("c++/executable")
	TypeCxxHeader         = Key("c++/header")
	TypeCxxLibrary        = Key("c++/library")
	TypeCxxLibraryStatic  = Key("c++/library/static")
	TypeCxxLibraryShared  = Key("c++/library/shared")
	TypeCxxLibraryPkgconf = Key("c++/library/pkgconf")
	TypeCxxLibraryDevlink = Key("c++/library/devlink")

	/* data files */
	TypeDataMisc    = Key("data/misc")
	TypeDataPixmap  = Key("data/pixmap")
	TypeDataDesktop = Key("data/desktop")

	/* locales */
	TypeI18nPo      = Key("i18n/po")
	TypeI18nDesktop = Key("i18n/desktop")

	/* documentation */
	TypeDocMan  = Key("doc/man")
	TypeDocMisc = Key("doc/misc")
	TypeDocYelp = Key("doc/yelp")

	/* generators */
	TypeGenGlibResource = Key("gen/glib-resource")
	TypeXdtCSource      = Key("gen/xdt-csource")
	TypeXxdCSource      = Key("gen/xxd-csource")
)
