package systempit

import (
	"fmt"
	"plumpit/protos"
	"time"

	proc "github.com/shirou/gopsutil/process"
)

type ProcPsutilSource struct{}

func (pSource ProcPsutilSource) GetProcCpuPercent(pid int32, duration time.Duration) (protos.PitMessage, error) {
	result := protos.PitMessage{}
	process := proc.Process{Pid: pid}
	exist, _ := proc.PidExists(pid)
	if !exist {
		return result, nil
	}
	percent, err := process.Percent(duration)
	if err != nil {
		fmt.Println("error !!!")
		return result, err
	}
	cpuResult := protos.ProcCpuPercent{
		Pid:     pid,
		Percent: float64(percent),
	}
	result = protos.PitMessage{
		PitType: protos.EnumPitType_PROC_CPU_PERCENT,
		Message: &protos.PitMessage_ProcCpuPercent{ProcCpuPercent: &cpuResult},
	}
	return result, err
}

func (self ProcPsutilSource) GetProcMemInfo(pid int32) (protos.PitMessage, error) {
	process := proc.Process{Pid: pid}
	mem, err := process.MemoryInfo()
	if err != nil {
		return protos.PitMessage{}, err
	}
	memResult := protos.ProcMemInfo{
		Pid:  pid,
		RSS:  mem.RSS,
		VMS:  mem.VMS,
		Swap: mem.Swap,
	}
	return protos.PitMessage{
		PitType: protos.EnumPitType_PROC_MEM_INFO,
		Message: &protos.PitMessage_ProcMemInfo{ProcMemInfo: &memResult},
	}, nil
}

func (pSource ProcPsutilSource) GetProcMemPercent(pid int32) (protos.PitMessage, error) {
	result := protos.PitMessage{}
	process := proc.Process{Pid: pid}
	exist, _ := proc.PidExists(process.Pid)
	if !exist {
		return result, nil
	}
	percent, err := process.MemoryPercent()
	if err != nil {
		return result, err
	}
	memResult := protos.ProcMemPercent{
		Pid:     pid,
		Percent: float64(percent),
	}
	result = protos.PitMessage{
		PitType: protos.EnumPitType_PROC_MEM_PERCENT,
		Message: &protos.PitMessage_ProcMemPercent{ProcMemPercent: &memResult},
	}
	return result, err
}
