package base

type BinaryFormat string

const (
	BinaryFormatELF = BinaryFormat("ELF")
	BiaryFormatMZ   = BinaryFormat("MZ")
)

type BinaryType string

const (
	BinaryTypeExecutable    = BinaryType("executable")
	BinaryTypeLibraryShared = BinaryType("library-shared")
	BinaryTypeLibraryStatic = BinaryType("library-static")
	BinaryTypeObject        = BinaryType("object")
)

type BinaryFileInfo struct {
	Filename   string
	Type       BinaryType
	Format     BinaryFormat
	Depends    []string
	BinaryArch string // potentially format-specific naming
}

func (b BinaryFileInfo) DependsInfo() string {
	s := ""
	for _, d := range b.Depends {
		s = s + d + " " + b.BinaryArch + "\n"
	}
	return s
}
