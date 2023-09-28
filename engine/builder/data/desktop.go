package data

import (
	"github.com/metux/go-metabuild/engine/builder/base"
	"github.com/metux/go-metabuild/spec"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/fileutil"
)

type BuilderDataDesktop struct {
	base.BaseBuilder
}

func (b BuilderDataDesktop) JobRun() error {
	text := []string{
		"[Desktop Entry]",
		"Version=1.1",
		"Type=" + b.RequiredEntryStr(target.KeyDesktopType),
		"Name=" + b.RequiredEntryStr(target.KeyDesktopName),
		"GenericName=" + b.EntryStr(target.KeyDesktopGenericName),
		"Comment=" + util.StrEscLF(b.EntryStr(target.KeyDesktopComment)),
		"Icon=" + b.EntryStr(target.KeyDesktopIconFile),
		"Exec=" + b.EntryStr(target.KeyDesktopExec),
		"TryExec=" + b.EntryStr(target.KeyDesktopTryExec),
		"Terminal=" + b.EntryStr(target.KeyDesktopTerminal),
		"Categories=" + b.EntryStr(target.KeyDesktopCategories)}

	if err := fileutil.WriteFileLines(b.OutputFile(), text); err != nil {
		return err
	}

	b.InstallPkgFileAuto()
	return nil
}

// FIXME: move this to base builder ?
func (b BuilderDataDesktop) JobPrepare(id string) error {
	b.LoadTargetDefaults(spec.Key(b.Type()))
	return nil
}

func MakeDataDesktop(o spec.TargetObject, id string) BuilderDataDesktop {
	return BuilderDataDesktop{base.BaseBuilder{o, id}}
}
