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

	/* generated C */
	TypeCGlibMarshal  = Key("c/glib-marshal")
	TypeCGlibResource = Key("c/glib-resource")

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
	TypeI18nPo     = Key("i18n/po")
	TypeI18nMsgfmt = Key("i18n/msgfmt")

	/* documentation */
	TypeDocMan  = Key("doc/man")
	TypeDocMisc = Key("doc/misc")
	TypeDocYelp = Key("doc/yelp")

	/* generators */
	TypeGenGlibResource = Key("gen/glib-resource")
	TypeGenGlibMarshal  = Key("gen/glib-marshal")
	TypeXdtCSource      = Key("gen/xdt-csource")
	TypeXxdCSource      = Key("gen/xxd-csource")

	/* shell exec */
	TypeExecBasic = Key("exec/basic")
)
