package systempit

import (
	"plumpit/protos"

	"github.com/cloudfoundry/gosigar"
)

type SigarSource struct{}

func (self SigarSource) GetSystemCpu() protos.PitMessage {
	var cpu sigar.Cpu
	cpu.Get()
	cpuRes := protos.SystemCpu{
		PitType: protos.EnumPitType_SYSTEM_CPU,
		Sys:     cpu.Sys,
		Idle:    cpu.Idle,
		User:    cpu.User,
	}

	return protos.PitMessage{
		Message: &protos.PitMessage_SystemCpu{
			SystemCpu: &cpuRes,
		},
	}
}
