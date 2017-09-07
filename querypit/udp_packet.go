package querypit

import (
	"fmt"
	"plumpit/base"
	"plumpit/protos"
	"time"
)

type NameData [64]byte

const (
	GpmonPktTypeNone = iota
	GpmonPktTypeHello
	GpmonPktTypeMetrics
	GpmonPktTypeQlog
	GpmonPktTypeQexec
)

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
type GpmonNodeKey struct {
	SegID int16
	Dummy int16
	PID   int32
	NID   int32
}
type GpmonQexecKey struct {
	QKey GpmonQlogKey
	NKey GpmonNodeKey
}

func (k GpmonNodeKey) GetHashKeyString() string {
	return fmt.Sprintf("%d-%d-%d", k.SegID, k.PID, k.NID)
}

// GpmonQexec is the mirror type for gpmon_qexec_t in GPDB.
type GpmonQexec struct {
	Key                    GpmonQexecKey
	Hname                  NameData
	Status                 uint64
	CPUElapsed             uint64
	PMetrics               GpmonProcMetrics
	Rowsout, Rowsin        uint64
	StartupCost, TotalCost float64
	PlanRows               float64
	NodeType               int32
	Dummy2                 [4]byte
	Offset                 int64
}

func (q GpmonQlog) ToPitMessage() (protos.PitMessage, error) {
	if gShmSourceMap[q.SharedMemory] == nil {
		shmSource := &GPShmSource{
			ShmID:             q.SharedMemory,
			InstrumentOffsets: map[string]int64{},
		}
		gShmSourceMap[q.SharedMemory] = shmSource
		go shmSource.Run(time.Duration(5))
	}
	gRunningQueriesShmInfo[GetQueryIdString(q.Key)] = q.SharedMemory
	query := protos.QueryInfo{}
	// Do content transfer here
	return protos.PitMessage{
		PitType: protos.EnumPitType_QUERY_INFO,
		Message: &protos.PitMessage_QueryInfo{
			QueryInfo: &query,
		},
	}, nil
}
func GetQueryIdString(k GpmonQlogKey) string {
	return fmt.Sprintf("%d-%d-%d", k.Tmid, k.Ssid, k.Ccnt)
}
func (q GpmonQexec) ToPitMessage() (protos.PitMessage, error) {
	queryid := GetQueryIdString(q.Key.QKey)
	key := protos.DistributedNodeKey{
		SegId:  int32(q.Key.NKey.SegID),
		ProcId: q.Key.NKey.PID,
		NodeId: q.Key.NKey.NID,
	}
	if q.Offset != 0 {
		if shmid, ok := gRunningQueriesShmInfo[queryid]; ok {
			shmSource := gShmSourceMap[shmid]

			shmSource.InstrumentOffsets[q.Key.NKey.GetHashKeyString()] = q.Offset
		}
	}
	exec := protos.ExecInfo{
		QueryId:     queryid,
		NodeKey:     &key,
		Status:      protos.EnumNodeStatus(q.Status),
		PlanRows:    q.PlanRows,
		NodeType:    q.NodeType,
		StartupCost: q.StartupCost,
		TotalCost:   q.TotalCost,
	}
	return protos.PitMessage{
		PitType: protos.EnumPitType_EXEC_INFO,
		Message: &protos.PitMessage_ExecInfo{
			ExecInfo: &exec,
		},
	}, nil
}
