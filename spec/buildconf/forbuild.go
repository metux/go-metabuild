package buildconf

import (
	"github.com/metux/go-metabuild/util"
)

// Keys for target branch subtree underneath buildconf
const (
	KeyForBuild = Key("build")
	KeyForHost  = Key("host")
)

// buildconf has subtrees for items specific to whether we're building for the
// (currently running) *build* machine or the final target *host*
// SubForBuild() retrieves the corresponding subtree as a SpecObj
// build=true means fetching for the currently building machine, otherwise the
// the target host
func (bc BuildConf) SubForBuild(build bool, k Key) SpecObj {
	return bc.EntrySpec(util.ValIf(build, KeyForBuild, KeyForHost).Append(k))
}
