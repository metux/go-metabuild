package base

import (
	"testing"
)

var testdata = map[string]Machine{
	"x86_64-unknown-linux": Machine{
		Arch:   "x86_64",
		Kernel: "linux",
		Vendor: "unknown",
	},
	"x86_64-unknown-linux-gnu": Machine{
		Arch:   "x86_64",
		Kernel: "linux",
		Vendor: "unknown",
		System: "gnu",
	},
	"x86_64-gdcproject-linux-gnu": Machine{
		Arch:   "x86_64",
		Kernel: "linux",
		Vendor: "gdcproject",
		System: "gnu",
	},
	"arm-unknown-linux-gnueabi": Machine{
		Arch:   "arm",
		Kernel: "linux",
		Vendor: "unknown",
		System: "gnueabi",
	},
	"arm-unknown-linux-androideabi": Machine{
		Arch:   "arm",
		Kernel: "linux",
		Vendor: "unknown",
		System: "androideabi",
	},
	"x86_64-w64-mingw32": Machine{
		Arch:   "x86_64",
		Kernel: "mingw32",
		Vendor: "w64",
	},
	"x86_64-pc-mingw32": Machine{
		Arch:   "x86_64",
		Kernel: "mingw32",
		Vendor: "pc",
	},
	"i686-linux-gnu": Machine{
		Arch:   "i686",
		Kernel: "linux",
		Vendor: "",
		System: "gnu",
	},
}

func TestMachine(t *testing.T) {
	for str, ent := range testdata {
		mach := ParseMachine(str)
		if ent.Arch != mach.Arch {
			t.Errorf("[%s] arch mismatch: have %s <-> want %s\n", str, ent.Arch, mach.Arch)
		}
		if ent.Kernel != mach.Kernel {
			t.Errorf("[%s] kernel mismatch: have %s <-> want %s\n", str, ent.Kernel, mach.Kernel)
		}
		if ent.Vendor != mach.Vendor {
			t.Errorf("[%s] vendor mismatch: have %s <-> want %s\n", str, ent.Vendor, mach.Vendor)
		}
		if ent.System != mach.System {
			t.Errorf("[%s] system mismatch: have %s <-> want %s\n", str, ent.System, mach.System)
		}
		if str != ent.String() {
			t.Errorf("[%s] serialize mismatch\n", ent.String())
		}
		if str != mach.String() {
			t.Errorf("[%s] serialize mismatch\n", mach.String())
		}
	}
}
