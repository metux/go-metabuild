package check

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/segmentio/fasthash/fnv1a"

	"github.com/metux/go-metabuild/spec/buildconf"
	"github.com/metux/go-metabuild/spec/cache"
	"github.com/metux/go-metabuild/spec/distro"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/specobj"
)

type Check struct {
	specobj.SpecObj
	Cache     cache.Cache
	BuildConf buildconf.BuildConf
	Distros   distro.Distro
}

type CheckList = []Check

// Check whether result already cached and return it
// -> (cached, cacheval)
func (chk Check) Cached() (bool, bool) {
	return chk.Cache.CheckGet(chk.Hash())
}

func (chk Check) IsMandatory() bool {
	return chk.EntryBoolDef(KeyMandatory, false)
}

func (chk Check) IsDone() bool {
	return chk.EntryBoolDef(KeyDone, false)
}

func (chk Check) IsPresent() bool {
	return chk.EntryBoolDef(KeyPresent, false)
}

func (chk Check) ForBuild() bool {
	return chk.EntryBoolDef(KeyBuild, false)
}

func (chk Check) TempDir() string {
	return chk.BuildConf.BuildTempDir("checks/" + chk.Hash())
}

func (chk Check) Done(succ bool) {
	// FIXME: write to cache
	chk.EntryPutBool(KeyPresent, succ)
	chk.EntryPutBool(KeyDone, true)
}

// Compute a hash of the attributes that describing the check's tests,
// but skipping any other ones, eg. those telling what to do with the result
func (chk Check) Hash() string {
	h := fnv1a.Init64
	for _, x := range chk.EntryStrList(KeyHashAttrs) {
		h = fnv1a.AddString64(h, chk.EntryStr(Key(x)))
	}
	return fmt.Sprintf("%X", h)
}

func (chk Check) Init() error {
	if err := chk.detectCheckType(); err != nil {
		return err
	}
	return chk.initFields()
}

func (chk Check) String() string {
	h := chk.Type() + " "
	for _, k := range chk.EntryStrList(KeyHashAttrs) {
		if k != string(KeyType) {
			h = fmt.Sprintf("%s %s=%s ", h, k, chk.EntryStrList(Key(k)))
		}
	}
	return h + util.ValIf(chk.IsMandatory(), " (mandatory)", "") + " " + chk.Hash()
}

func (chk Check) TestProgFile(suffix string) string {
	return fmt.Sprintf("%s/metabuild-check-%s-%s.%s",
		os.TempDir(), util.ValIf(chk.ForBuild(), "build", "host"),
		chk.Hash(), suffix)
}

func (chk Check) GetDefines() []string {
	return chk.EntryStrList(KeyDefines)
}

func (chk Check) Type() string {
	return chk.EntryStr(KeyType)
}

func (chk Check) detectCheckType() error {
	if chk.Type() != "" {
		return nil
	}

	types := []Key{
		// c-header needs to be last, since others may also have headers
		KeyCFunction,
		KeyCType,
		KeyCHeader,

		KeyPkgConfig,
	}

	for _, x := range types {
		if s := chk.EntryStrList(x); len(s) > 0 {
			chk.EntryPutStr(KeyType, string(x))
			return nil
		}
	}

	return ErrUnsupportedCheck
}

func (chk Check) initFields() error {
	chk.EntryPutStrList(KeyHashAttrs, []string{
		string(KeyType),
		string(KeyCHeader),
		string(KeyCFunction),
		string(KeyCType),
		string(KeyPkgConfig),
		string(KeyBuild),
	})
	return nil
}

func (chk Check) Id() string {
	return chk.EntryStr(KeyId)
}

func (chk Check) SetId(id string) {
	chk.EntryPutStr(KeyId, id)
}

func (chk Check) SetIdList(ids []string) []string {
	chk.SetId(strings.Join(ids, " "))
	return ids
}

func (chk Check) YesNoStrList(yesno bool, k Key) []string {
	return chk.EntryStrList(Key(util.YesNo(yesno) + "/" + string(k)))
}

func (chk Check) Logf(f string, args ...any) {
	if id := chk.Id(); id == "" {
		log.Printf("[Check %s] %s\n", chk.Type(), fmt.Sprintf(f, args...))
	} else {
		log.Printf("[Check %s %s] %s\n", chk.Type(), id, fmt.Sprintf(f, args...))
	}
}
