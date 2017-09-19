package collator

import (
	"fmt"
	"plumpit/base"
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

// NewProcessCollator create a new ProcessCollator.
func NewProcessCollator() *ProcessCollator {
	coll := &ProcessCollator{procs: map[int32]ProcessInfo{}, lock: new(sync.RWMutex)}
	return coll
}
func (p *ProcessCollator) DebugString() string {
	result := ""
	for pid, proc := range p.procs {
		instrs := proc.info.GetInsight()
		if len(instrs) > 0 {
			for _, instr := range instrs {
				result += fmt.Sprintf("%d\tseg<%2d>\tnode<%2d>\t%d\n", pid, instr.Key.SegId, instr.Key.NodeId, instr.TupleCount)
			}
		}
	}
	return result
}

// ProcExist check if pid exist in the process map.
func (p *ProcessCollator) ProcExist(pid int32) bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	_, ok := p.procs[pid]
	return ok
}

// CreateProc add the pid in the process map
func (p *ProcessCollator) CreateProc(pid int32) {
	p.lock.Lock()
	defer p.lock.Unlock()
	_, ok := p.procs[pid]
	if !ok {
		p.procs[pid] = ProcessInfo{
			lock: new(sync.RWMutex),
			info: protos.ProcessSamplingInfo{
				Pid:     pid,
				Insight: map[int32]*protos.PlumInstrument{},
			},
		}
	}
}

func (p *ProcessCollator) CollateExecInfo(msg protos.PitMessage) {
	q := msg.GetExecInfo()
	pid := q.GetNodeKey().ProcId
	switch q.GetStatus() {
	case protos.EnumNodeStatus_NODE_INIT, protos.EnumNodeStatus_NODE_START:
		if !p.ProcExist(pid) {
			p.CreateProc(pid)
		}
	}
}

func (p *ProcessCollator) CollateInstrumentInfo(msg protos.PitMessage) {
	instr := msg.GetInstrument()
	pid := instr.GetKey().ProcId
	nid := instr.GetKey().NodeId
	if p.ProcExist(pid) {
		p.procs[pid].lock.Lock()
		defer p.procs[pid].lock.Unlock()
		p.procs[pid].info.Insight[nid] = instr
	}
}

func (p *ProcessCollator) Collate(msg protos.PitMessage) {
	switch msg.PitType {
	case protos.EnumPitType_EXEC_INFO:
		p.CollateExecInfo(msg)
	case protos.EnumPitType_INSTRUMENT_INFO:
		p.CollateInstrumentInfo(msg)
	case protos.EnumPitType_PROC_MEM_INFO:
	default:
	}
}

// AddRawMessage add a RawMessage and will call ToPitMessage then add it.
func (p *ProcessCollator) AddMessageFunc(sender base.Sender) func(base.RawMessage) error {
	return func(msg base.RawMessage) error {
		if msg == nil {
			return nil
		}
		pit, err := msg.ToPitMessage()
		if err != nil {
			return err
		}
		switch pit.PitType {
		case protos.EnumPitType_EXEC_INFO:
			p.Collate(pit)
			err = sender(pit)
		case protos.EnumPitType_QUERY_INFO:
			err = sender(pit)
		default:
			p.Collate(pit)
		}
		return err
	}

}
