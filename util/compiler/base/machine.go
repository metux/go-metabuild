package base

import (
	"fmt"
	"strings"
)

type Machine struct {
	Arch   string
	Vendor string
	Kernel string
	System string
	Valid  bool
}

func (mach Machine) String() string {
	if !mach.Valid {
		return "unknown"
	}
	if mach.Vendor == "" {
		if mach.System == "" {
			return fmt.Sprintf("%s-%s", mach.Arch, mach.Kernel)
		}
		return fmt.Sprintf("%s-%s-%s", mach.Arch, mach.Kernel, mach.System)
	}
	if mach.System == "" {
		return fmt.Sprintf("%s-%s-%s", mach.Arch, mach.Vendor, mach.Kernel)
	}
	return fmt.Sprintf("%s-%s-%s-%s", mach.Arch, mach.Vendor, mach.Kernel, mach.System)
}

func ParseMachine(mach string) Machine {
	s := strings.Split(strings.TrimSpace(mach), "-")
	if len(s) == 0 || s[0] == "" {
		return Machine{}
	}
	if len(s) == 3 {
		// no vendor field
		if s[1] == "linux" {
			return Machine{Arch: s[0], Kernel: s[1], System: s[2], Valid: true}
		}
		// FIXME: are there other special cases ?
		return Machine{Arch: s[0], Vendor: s[1], Kernel: s[2], Valid: true}
	}
	return Machine{Arch: s[0], Vendor: s[1], Kernel: s[2], System: s[3], Valid: true}
}
