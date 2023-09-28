package spec

import (
	"github.com/metux/go-metabuild/util/specobj"

	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/spec/check"
	"github.com/metux/go-metabuild/spec/distro"
	"github.com/metux/go-metabuild/spec/generate"
	"github.com/metux/go-metabuild/spec/global"
	"github.com/metux/go-metabuild/spec/target"
)

type (
	Key = specobj.Key

	Global       = global.Global
	Distro       = distro.Distro
	DistPkg      = distro.DistPkg
	BuildConf    = buildconf.BuildConf
	Generate     = generate.Generate
	Check        = check.Check
	TargetObject = target.TargetObject
)
