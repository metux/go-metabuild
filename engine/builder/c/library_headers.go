package c

type BuilderCLibraryHeaders struct {
	CommonCBuilder
}

func (b BuilderCLibraryHeaders) JobRun() error {
	if b.WantInstall() {
		dir := b.InstallDir()
		mode := b.InstallPerm()

		for _, s := range b.Sources() {
			b.InstallPkgFile(s, dir, mode)
		}
	}
	return nil
}
