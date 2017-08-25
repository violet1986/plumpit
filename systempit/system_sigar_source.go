package systempit

import (
	"plumpit/protos"

	"github.com/cloudfoundry/gosigar"
)

type SigarSource struct{}

func (self SigarSource) GetSystemCpu() (protos.PitMessage, error) {
	var cpu sigar.Cpu
	err := cpu.Get()
	if err != nil {
		return protos.PitMessage{}, err
	}
	cpuRes := protos.SystemCpu{
		Sys:  cpu.Sys,
		Idle: cpu.Idle,
		User: cpu.User,
	}

	return protos.PitMessage{
		PitType: protos.EnumPitType_SYSTEM_CPU,
		Message: &protos.PitMessage_SystemCpu{
			SystemCpu: &cpuRes,
		},
	}, nil
}
