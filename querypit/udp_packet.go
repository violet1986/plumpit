package querypit

import (
	"fmt"
	"plumpit/base"
	"plumpit/protos"
)

type NameData [64]byte

type GpmonPacket struct {
	Magic   int32
	Version int16
	Pkttype int16
}
type GpmonQlogKey struct {
	Tmid int32
	Ssid int32
	Ccnt int32
}

type GpmonProcMetrics struct {
	FdCnt                    uint32
	CPUPct                   float32
	MemSize, Resident, Share uint64
}

// GpmonQlog is the mirror type for gpmon_qlog_t in GPDB.
type GpmonQlog struct {
	Key                   GpmonQlogKey
	User                  NameData
	Database              NameData
	Tsubmit, Tstart, Tfin int32
	Status                int32
	Cost                  int32
	CPUElapsed            int64
	PMetrics              GpmonProcMetrics
	SharedMemory          base.SharedMemoryInfo
}

type GpmonQExecKey struct {
	QKey  GpmonQlogKey
	SegID int16
	Dummy int16
	PID   int32
	NID   int32
}

// GpmonQExec is the mirror type for gpmon_qexec_t in GPDB.
type GpmonQexec struct {
	Key        GpmonQExecKey
	Hname      NameData
	Status     uint16
	CPUElapsed uint64
	PMetrics   GpmonProcMetrics
	// TODO find out where should add this Dummy
	Dummy                  [6]byte
	Rowsout, Rowsin        uint64
	StartupCost, TotalCost float64
	PlanRows               float64
	NodeType               int32
	Offset                 uint64
}

func (q GpmonQlog) ToPitMessage() (protos.PitMessage, error) {
	query := protos.QueryInfo{}
	// Do content transfer here
	fmt.Println("shmid is", q.SharedMemory)
	return protos.PitMessage{
		PitType: protos.EnumPitType_QUERY_INFO,
		Message: &protos.PitMessage_QueryInfo{
			QueryInfo: &query,
		},
	}, nil
}

func (q GpmonQexec) ToPitMessage() (protos.PitMessage, error) {
	exec := protos.ExecInfo{}
	fmt.Println(&q.Key.QKey.Tmid, q.Rowsout, q.StartupCost, &q.StartupCost, q.PlanRows, &q.PlanRows, q.NodeType, &q.NodeType, q.Offset, &q.Offset)
	return protos.PitMessage{
		PitType: protos.EnumPitType_EXEC_INFO,
		Message: &protos.PitMessage_ExecInfo{
			ExecInfo: &exec,
		},
	}, nil
}
