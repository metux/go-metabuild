package c

import (
	"strings"

	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/compiler"
)

type BuilderCLibraryPkgConfig struct {
	CommonCBuilder
}

func (b BuilderCLibraryPkgConfig) JobRun() error {
	libname := b.RequiredEntryStr(target.KeyPkgLibname)

	requires := []string{}
	for _, pi := range b.Parent.AllImports() {
		if !pi.Private {
			requires = append(requires, pi.PkgSpec)
		}
	}

	pkginfo := compiler.PkgConfigInfo{
		Name:          b.RequiredEntryStr(target.KeyPkgName),
		Version:       b.RequiredEntryStr(target.KeyPkgVersion),
		Description:   b.RequiredEntryStr(target.KeyPkgDescription),
		SharedLdflags: []string{"-L${sharedlibdir}", "-l" + libname},
		SharedCflags:  []string{"-I${includedir}"}, /* add cflags ? */
		StaticLdflags: []string{"-L${libdir}", "-l" + libname},
		StaticCflags:  []string{"-I${includedir}"}, /* add cflags ? */
		Requires:      requires,
		Variables: map[string]string{
			"prefix":       b.RequiredEntryStr(target.KeyPkgPrefix),
			"exec_prefix":  b.RequiredEntryStr(target.KeyPkgExecPrefix),
			"libdir":       b.RequiredEntryStr(target.KeyPkgLibdir),
			"includedir":   b.RequiredEntryStr(target.KeyPkgIncludedir),
			"sharedlibdir": b.RequiredEntryStr(target.KeyPkgSharedLibdir),
			"archive":      b.RequiredEntryStr(target.KeyPkgArchive),
		},
	}

	pcfile := b.RequiredEntryStr(target.KeyFile)

	if err := pkginfo.Write(pcfile); err != nil {
		return err
	}

	if b.InstallPkgFileAuto() {
		// write pkgconf dependencies
		pkgdeps := []string{}
		for _, x := range requires {
			// we may have either "<name>" or "<name> >= <ver>"
			f := strings.Fields(x)
			if len(f) < 1 {
				b.Logf("WARN: empty pkgconfig req line\n")
			} else if len(f) < 3 {
				pkgdeps = append(pkgdeps, f[0])
			} else {
				pkgdeps = append(pkgdeps, f[0]+" "+f[2])
			}
		}
		b.WritePkgMeta(pcfile+".pcdep", strings.Join(pkgdeps, "\n"))
	}

	return nil
}
