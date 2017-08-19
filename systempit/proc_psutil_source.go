package systempit

import (
	"fmt"
	"log"
	"plumpit/protos"

	proc "github.com/shirou/gopsutil/process"
)

type ProcPsutilSource struct {
	process proc.Process
}

func NewProcPsutilSource(pid int) *ProcPsutilSource {
	process := proc.Process{Pid: int32(pid)}
	return &ProcPsutilSource{process: process}
}
func (self ProcPsutilSource) GetProcCPU() protos.ProcCPU {
	fmt.Print("self.process.Pid is ", self.process.Pid)
	percent, err := self.process.Percent(0)
	if err != nil {
		log.Println("Error during GetProcCPU:", err)
	}
	return protos.ProcCPU{
		PitType: protos.EnumPitType_PROC_CPU,
		Percent: percent,
	}

}
