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

func (bc BuildConf) BuildTempDir(sub string) string {
	d := BuildDirTmp + sub
	os.MkdirAll(d, 0755)
	return filepath.Clean(d)
}

func (bc BuildConf) BuildDistDir(sub string) string {
	d := BuildDirDist + sub
	os.MkdirAll(d, 0755)
	return filepath.Clean(d)
}

func (bc BuildConf) BuildDistPkgDir(pkg string) string {
	d := BuildDirDist + bc.EntryStr(KeyTargetDistName) + "/" + pkg
	os.MkdirAll(d, 0755)
	return filepath.Clean(d)
}

func (bc BuildConf) BuildDistPkgMetaDir(pkg string) string {
	d := BuildDirDist + bc.EntryStr(KeyTargetDistName) + "/" + pkg + "/meta/"
	os.MkdirAll(d, 0755)
	return filepath.Clean(d)
}

func (bc BuildConf) BuildDistPkgRootDir(pkg string, sub string) string {
	d := BuildDirDist + bc.EntryStr(KeyTargetDistName) + "/" + pkg + "/root/" + sub
	os.MkdirAll(d, 0755)
	return filepath.Clean(d)
}

func (bc BuildConf) CleanBuildDir() {
	os.RemoveAll(BuildDir)
}
