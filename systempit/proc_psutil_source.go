package systempit

import (
	"plumpit/protos"
	"time"

	proc "github.com/shirou/gopsutil/process"
)

type ProcPsutilSource struct {
	process proc.Process
}

func NewProcPsutilSource(pid int) *ProcPsutilSource {
	process := proc.Process{Pid: int32(pid)}
	return &ProcPsutilSource{process: process}
}
func (self ProcPsutilSource) GetProcCpuPercent(duration interface{}) (protos.PitMessage, error) {
	result := protos.PitMessage{}
	exist, _ := proc.PidExists(self.process.Pid)
	if !exist {
		return result, nil
	}
	percent, err := self.process.Percent(duration.(time.Duration))
	if err != nil {
		return result, err
	}
	cpuResult := protos.ProcCpuPercent{
		Pid:     self.process.Pid,
		Percent: float64(percent),
	}
	result = protos.PitMessage{
		PitType: protos.EnumPitType_PROC_CPU_PERCENT,
		Message: &protos.PitMessage_ProcCpuPercent{ProcCpuPercent: &cpuResult},
	}
	return result, err
}

func (self ProcPsutilSource) GetProcMemInfo() (protos.PitMessage, error) {
	mem, err := self.process.MemoryInfo()
	if err != nil {
		return protos.PitMessage{}, err
	}
	memResult := protos.ProcMemInfo{
		Pid:  self.process.Pid,
		RSS:  mem.RSS,
		VMS:  mem.VMS,
		Swap: mem.Swap,
	}
	return protos.PitMessage{
		PitType: protos.EnumPitType_PROC_MEM_INFO,
		Message: &protos.PitMessage_ProcMemInfo{ProcMemInfo: &memResult},
	}, nil
}

func (self ProcPsutilSource) GetProcMemPercent() (protos.PitMessage, error) {
	result := protos.PitMessage{}
	exist, _ := proc.PidExists(self.process.Pid)
	if !exist {
		return result, nil
	}
	percent, err := self.process.MemoryPercent()
	if err != nil {
		return result, err
	}
	memResult := protos.ProcMemPercent{
		Pid:     self.process.Pid,
		Percent: float64(percent),
	}
	result = protos.PitMessage{
		PitType: protos.EnumPitType_PROC_MEM_PERCENT,
		Message: &protos.PitMessage_ProcMemPercent{ProcMemPercent: &memResult},
	}
	return result, err
}
