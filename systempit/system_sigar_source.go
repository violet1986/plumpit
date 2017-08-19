package systempit

import (
	"plumpit/protos"

	"github.com/cloudfoundry/gosigar"
)

type SigarSource struct{}

func (self SigarSource) GetSystemCPU() protos.SystemCPU {
	var cpu sigar.Cpu
	cpu.Get()
	return protos.SystemCPU{
		PitType: protos.EnumPitType_SYSTEM_CPU,
		Sys:     cpu.Sys,
		Idle:    cpu.Idle,
		User:    cpu.User,
	}
}
