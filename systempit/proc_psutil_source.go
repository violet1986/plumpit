package systempit

import (
	"log"
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
func (self ProcPsutilSource) GetProcCPU(duration interface{}) protos.ProcCPU {
	percent, err := self.process.Percent(duration.(time.Duration) * time.Second)
	if err != nil {
		log.Println("Error during GetProcCPU:", err)
	}
	return protos.ProcCPU{
		PitType: protos.EnumPitType_PROC_CPU,
		Percent: float64(percent),
	}
}
