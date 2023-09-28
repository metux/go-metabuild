package buildconf

const (
	KeyInstallDirs = Key("install-dirs")

	KeyInstallDirPrefix    = Key("prefix")
	KeyInstallDirData      = Key("datadir")
	KeyInstallDirBin       = Key("bindir")
	KeyInstallDirLib       = Key("libdir")
	KeyInstallDirPkgConfig = Key("pkgconfigdir")
	KeyInstallDirMan       = Key("mandir")
	KeyInstallDirDoc       = Key("docdir")
	KeyInstallDirInclude   = Key("includedir")

	KeyDestDir = Key("DESTDIR")
)

func (bc BuildConf) InstallDirs() SpecObj {
	return bc.EntrySpec(KeyInstallDirs)
}
