package base

type CompilerArg struct {
	IncludeDirsAbs     []string
	IncludeDirsSysroot []string
	Sources            []string
	Output             string
	DllName            string
	Defines            []string
	PkgImports         []PkgConfigInfo
}
