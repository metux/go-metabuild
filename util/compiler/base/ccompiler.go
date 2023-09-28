package base

type CCompiler interface {
	CompileExecutable(CompilerArg) error
	CompileLibraryStatic(CompilerArg) error
	CompileLibraryShared(CompilerArg) error
	BinaryInfo(string) BinaryFileInfo
}
