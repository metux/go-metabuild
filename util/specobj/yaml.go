package specobj

import (
	"os"

	magic "github.com/metux/go-magicdict"
)

func (so SpecObj) YamlStore(fn string, fmode os.FileMode) error {
	return magic.YamlStore(fn, so.Spec, fmode)
}

func YamlLoad(fn string, dflt string) (SpecObj, error) {
	md, err := magic.YamlLoad(fn, dflt)
	return SpecObj{md, err}, err
}
