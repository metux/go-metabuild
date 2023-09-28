package engine

type Stage string

const (
	StageInit      = Stage("init")
	StageConfigure = Stage("configure")
	StageAutogen   = Stage("autogen")
	StageBuild     = Stage("build")
	StagePackage   = Stage("package")
)
