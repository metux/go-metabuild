package dpkg

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/fileutil"
	"github.com/metux/go-metabuild/util/pkgmgr/base"
)

type Dpkg struct {
	Command []string
	Env     []string
	Sysroot string
}

func NewDpkg(command []string, env []string, sysroot string) Dpkg {
	if len(command) == 0 {
		command = []string{"dpkg"}
	}
	return Dpkg{command, env, sysroot}
}

func (dpkg Dpkg) SearchInstalledFile(fn string) []base.PkgFileEntry {
	res := make([]base.PkgFileEntry, 0)

	cmdline := append(dpkg.Command, "-S", fn)
	out, err := cmd.RunOutLines(cmdline, true)

	if err != nil {
		return res
	}

	for _, line := range out {
		a := strings.Split(line, ": ")
		res = append(res, base.NewPkgFileEntry(a[0], a[1]))
	}

	return res
}

func (dpkg Dpkg) MatchElfArch(debArch string, elfArch string) bool {
	if debArch == elfArch {
		return true
	}
	// FIXME: need a better mapping for that
	if strings.HasSuffix(elfArch, debArch) {
		return true
	}

	if debArch == "amd64" && elfArch == "elf64-x86-64" {
		return true
	}

	return false
}

func (dpkg Dpkg) WriteControlFile(pkgroot string, control base.PkgControl) error {
	debroot := pkgroot + "/DEBIAN"
	os.MkdirAll(debroot, 0755)

	// write DEBIAN/control
	wr, err := fileutil.CreateWriter(debroot + "/control")
	defer wr.Close()
	if err != nil {
		return err
	}

	wr.WriteFolded("Package", control.Package)
	wr.WriteFolded("Version", control.Version)
	wr.WriteFolded("Maintainer", control.Maintainer)
	wr.WriteFolded("Section", control.Section)
	wr.WriteFolded("Depend", strings.Join(control.Depend, ",\n"))
	wr.WriteFolded("Multi-Arch", control.MultiArch)
	wr.WriteFolded("Bugs", control.Bugs)
	wr.WriteFolded("Homepage", control.Homepage)
	wr.WriteFolded("Origin", control.Origin)
	wr.WriteFolded("Priority", control.Priority)

	if control.Architecture == "" {
		wr.WriteFolded("Architecture", "all")
	} else {
		wr.WriteFolded("Architecture", control.Architecture)
	}

	wr.WriteFolded("Description", control.Description)

	// write DEBIAN/md5sums
	md5, err := fileutil.MD5Files(pkgroot, dpkg.ScanPlainFiles(pkgroot))
	if err != nil {
		return err
	}
	return fileutil.WriteText(debroot+"/md5sums", md5)
}

func (dpkg Dpkg) ScanPlainFiles(rootdir string) []string {
	rootdirlen := len(rootdir)
	flist := []string{}
	filepath.Walk(rootdir,
		func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && (info.Mode()&os.ModeSymlink != os.ModeSymlink) {
				p := path[rootdirlen:]
				if !strings.HasPrefix(p, "DEBIAN/") {
					flist = append(flist, p)
				}
			}
			return err
		})
	return flist
}

func (dpkg Dpkg) Build(pkgroot string, targetdir string) error {
	cmdline := []string{"dpkg-deb", "-b", "--root-owner-group", pkgroot, targetdir}
	_, err := cmd.RunOutLines(cmdline, true)
	if err != nil {
		log.Println("DPKG: build error", err)
		return err
	}
	return nil
}
