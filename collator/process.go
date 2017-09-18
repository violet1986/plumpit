package collator

import (
	"plumpit/protos"
	"sync"
)

// ProcessInfo represents a process's information, including system metrics and query metrics.
type ProcessInfo struct {
	info protos.ProcessSamplingInfo
	lock *sync.RWMutex
}

// ProcessCollator contains all processes' information in one host
type ProcessCollator struct {
	RawCollator
	procs map[int32]ProcessInfo
	lock  *sync.RWMutex
}

func (p *ProcessCollator) ProcExist(pid int32) bool {
	p.lock.RLock()
	_, ok := p.procs[pid]
	p.lock.RUnlock()
	return ok
}
func (p *ProcessCollator) CreateProc(pid int32) {
	p.lock.Lock()
	_, ok := p.procs[pid]
	if ok {
		return
	}
	p.procs[pid] = ProcessInfo{lock: new(sync.RWMutex)}
	p.lock.Unlock()
}
func (p *ProcessCollator) AddMessage(msg protos.PitMessage) {
	switch msg.PitType {
	case protos.EnumPitType_EXEC_INFO:
		q := msg.GetExecInfo()
		pid := q.GetNodeKey().ProcId
		switch q.GetStatus() {
		case protos.EnumNodeStatus_NODE_INIT:
		case protos.EnumNodeStatus_NODE_START:
			if !p.ProcExist(pid) {
				p.CreateProc(pid)
			}
		}
		p.RawCollator.AddMessage(msg)
	case protos.EnumPitType_INSTRUMENT_INFO:
		instr := msg.GetInstrument()
		pid := instr.GetKey().ProcId
		nid := instr.GetKey().NodeId
		if p.ProcExist(pid) {
			p.procs[pid].lock.Lock()
			p.procs[pid].info.Insight[nid] = instr
			p.procs[pid].lock.Unlock()
		}
	case protos.EnumPitType_PROC_MEM_INFO:
	default:
		p.RawCollator.AddMessage(msg)
	}

}
