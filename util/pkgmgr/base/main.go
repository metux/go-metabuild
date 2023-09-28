package base

type PackageManager interface {
	SearchInstalledFile(fn string) []PkgFileEntry
	MatchElfArch(debArch string, elfArch string) bool
	WriteControlFile(fn string, ctrl PkgControl) error
	Build(pkgroot string, targetdir string) error
}
