package engine

import (
	"log"
	"os"

	"github.com/metux/go-metabuild/engine/autoconf"
	"github.com/metux/go-metabuild/engine/builder"
	"github.com/metux/go-metabuild/engine/packager"
	"github.com/metux/go-metabuild/spec/global"
	"github.com/metux/go-metabuild/util"
)

var (
	ErrNoStage = util.Error("no such stage")
)

type Engine struct {
	Global global.Global
}

func (e *Engine) Load(fn string, dflt string) error {
	if gs, err := global.LoadGlobal(fn, dflt); err == nil {
		e.Global = gs
	} else {
		return err
	}
	return nil
}

func (e Engine) StageInit() error {
	log.Printf("Stage: init: %s ver %s\n",
		e.Global.EntryStr(global.KeyPackage),
		e.Global.EntryStr(global.KeyVersion))
	if srcdir := e.Global.EntryStr(global.KeySrcDir); srcdir != "" {
		if err := os.Chdir(srcdir); err != nil {
			return err
		}
		// prevent accidential double chdir
		e.Global.EntryDelete(global.KeySrcDir)
	}
	return nil
}

func (e Engine) StageConfigure() error {
	if err := e.StageInit(); err != nil {
		return err
	}
	log.Println("Stage: configure")
	return autoconf.RunConfigure(e.Global)
}

func (e Engine) StageAutogen() error {
	if err := e.StageConfigure(); err != nil {
		return err
	}
	log.Println("Stage: autogen")
	return autoconf.RunGenerates(e.Global)
}

func (e Engine) StageBuild() error {
	if err := e.StageAutogen(); err != nil {
		return err
	}
	log.Println("Stage: build")
	return builder.RunBuild(e.Global)
}

func (e Engine) StagePackage() error {
	if err := e.StageBuild(); err != nil {
		return err
	}
	log.Println("Stage: package")
	return packager.RunPackaging(e.Global)
}

func (e Engine) RunStage(s Stage) error {
	switch s {
	case StageInit:
		return e.StageInit()
	case StageConfigure:
		return e.StageConfigure()
	case StageAutogen:
		return e.StageAutogen()
	case StageBuild:
		return e.StageBuild()
	case StagePackage:
		return e.StagePackage()
	}
	return ErrNoStage
}

func Load(fn string, defaults string) (Engine, error) {
	e := Engine{}
	err := e.Load(fn, defaults)
	return e, err
}
