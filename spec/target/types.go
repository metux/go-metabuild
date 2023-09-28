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

	/* data files */
	TypeDataMisc    = Key("data/misc")
	TypeDataPixmaps = Key("data/pixmaps")
	TypeDataDesktop = Key("data/desktop")

	/* locales */
	TypeI18nPo = Key("i18n/po")

	/* documentation */
	TypeDocMan  = Key("doc/man")
	TypeDocMisc = Key("doc/misc")
)