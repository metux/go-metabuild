package buildconf

import (
	"os"
	"path/filepath"
)

const (
	BuildDir     = "./.build/"
	BuildDirTmp  = BuildDir + "tmp/"
	BuildDirDist = BuildDir + "dist/"
)

func xmkdir(d string) string {
	os.MkdirAll(d, 0755)
	d, _ = filepath.Abs(filepath.Clean(d))
	return d
}

func (bc BuildConf) BuildTempDir(sub string) string {
	return xmkdir(BuildDirTmp + sub)
}

func (bc BuildConf) BuildDistDir(sub string) string {
	return xmkdir(BuildDirDist + sub)
}

func (bc BuildConf) BuildDistPkgDir(pkg string) string {
	return xmkdir(BuildDirDist + bc.EntryStr(KeyTargetDistName) + "/" + pkg)
}

func (bc BuildConf) BuildDistPkgMetaDir(pkg string) string {
	return xmkdir(BuildDirDist + bc.EntryStr(KeyTargetDistName) + "/" + pkg + "/meta/")
}

func (bc BuildConf) BuildDistPkgRootDir(pkg string, sub string) string {
	return xmkdir(BuildDirDist + bc.EntryStr(KeyTargetDistName) + "/" + pkg + "/root/" + sub)
}

func (bc BuildConf) CleanBuildDir() {
	os.RemoveAll(BuildDir)
}
