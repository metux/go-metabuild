package buildconf

import (
	"github.com/metux/go-metabuild/spec/distro"
)

// Keys underneath buildconf::
const (
	KeyTargetDist     = Key("@targetdist")
	KeyTargetDistName = Key("@targetdistname")
	KeyTargetDistArch = Key("@targetdistarch")
	KeyTargetPlatform = Key("@target-platform")
)

// store the probed target system
// FIXME: need to differenciate between host and build ?
func (bc BuildConf) SetTargetDist(dist string) {
	bc.EntryPutStr(KeyTargetDistName, dist)
	bc.EntryPutStr(KeyTargetDist, "${distro::"+dist+"}")
	bc.EntryPutStr(KeyTargetPlatform, "${distro::"+dist+"::platform}")
}

func (bc BuildConf) SetTargetDistArch(arch string) {
	bc.EntryPutStr(KeyTargetDistArch, arch)
}

// retrieve spec of the current target distro
func (bc BuildConf) TargetDistro() distro.Distro {
	// explicitly giving the target dist name, since @@KEY won't work when using aliases
	return distro.NewDistro(bc.EntrySpec(KeyTargetDist), bc.EntryStr(KeyTargetDistName))
}
