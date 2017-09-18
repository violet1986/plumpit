package collator

import (
	"plumpit/protos"
	"sync"
)

type ProcessInfo struct {
	info protos.ProcessSamplingInfo
	lock *sync.RWMutex
}
type ProcessCollator struct {
	procs map[int32]ProcessInfo
}

func (p *ProcessCollator) AddMessage(msg protos.PitMessage) {
	switch msg.PitType {
	case protos.EnumPitType_INSTRUMENT_INFO:
		instr := msg.GetInstrument()
		pid := instr.GetKey().ProcId
		nid := instr.GetKey().NodeId
		p.procs[pid].lock.Lock()
		p.procs[pid].info.Insight[nid] = instr
		p.procs[pid].lock.Unlock()
	}
}
