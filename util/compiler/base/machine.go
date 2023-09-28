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
}

func (mach Machine) String() string {
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

func (m *Machine) Parse(mach string) {
	s := strings.Split(strings.TrimSpace(mach), "-")
	if len(s) == 3 {
		// no vendor field
		if s[1] == "linux" {
			*m = Machine{Arch: s[0], Kernel: s[1], System: s[2]}
		} else {
			// FIXME: are there other special cases ?
			*m = Machine{Arch: s[0], Vendor: s[1], Kernel: s[2]}
		}
	} else {
		*m = Machine{Arch: s[0], Vendor: s[1], Kernel: s[2], System: s[3]}
	}
}

func ParseMachine(mach string) Machine {
	s := strings.Split(strings.TrimSpace(mach), "-")
	if len(s) == 3 {
		// no vendor field
		if s[1] == "linux" {
			return Machine{Arch: s[0], Kernel: s[1], System: s[2]}
		}
		// FIXME: are there other special cases ?
		return Machine{Arch: s[0], Vendor: s[1], Kernel: s[2]}
	}
	return Machine{Arch: s[0], Vendor: s[1], Kernel: s[2], System: s[3]}
}
