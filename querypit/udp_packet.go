package querypit

import (
	"fmt"
	"plumpit/protos"
)

// NameData is for the char[] type in C.
type NameData [64]byte

const (
	gpmonPktTypeNone = iota
	gpmonPktTypeHello
	gpmonPktTypeMetrics
	gpmonPktTypeQlog
	gpmonPktTypeQexec
	gpmonPktTypeSegInfo
	gpmonPktTypeStat = 10
)

// GpmonPacket is the head part of all UDP packet.
type GpmonPacket struct {
	Magic   int32
	Version int16
	Pkttype int16
}

// GpmonQlogKey represents the query id.
type GpmonQlogKey struct {
	Tmid int32
	Ssid int32
	Ccnt int32
}

func (k GpmonQlogKey) GetQueryIdString() string {
	return fmt.Sprintf("%d-%d-%d", k.Tmid, k.Ssid, k.Ccnt)
}

// GpmonProcMetrics represents process metrics, it is not used now and might be deprecated later
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
}

// GpmonNodeKey represents the hash key of a exec plan node.
// Dummy is only for c & golang binary transfer.
type GpmonNodeKey struct {
	SegID int16
	Dummy int16
	PID   int32
	NID   int32
}

// GpmonQexecKey is a combination of Qlog key and node key.
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
}

func (q *GpmonQlog) ToPitMessage() (protos.PitMessage, error) {
	return protos.PitMessage{
		PitType: protos.EnumPitType_QUERY_INFO,
		Message: &protos.PitMessage_QueryInfo{
			QueryInfo: &protos.QueryInfo{
				QueryId: q.Key.GetQueryIdString(),
				Status:  protos.EnumQueryStatus(q.Status),
			},
		},
	}, nil
}
func (q *GpmonQexec) ToPitMessage() (protos.PitMessage, error) {
	queryid := q.Key.QKey.GetQueryIdString()
	key := protos.DistributedNodeKey{
		SegId:  int32(q.Key.NKey.SegID),
		ProcId: q.Key.NKey.PID,
		NodeId: q.Key.NKey.NID,
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

// GpmonStat is the instrument info. It may be replaced by complete instrument later.
type GpmonStat struct {
	Running    bool
	Dummy      byte
	SegID      int16
	PID        int32
	NID        int32
	Dummy2     [4]byte
	TupleCount uint64
	NTuples    uint64
	NLoops     uint64
	FirstTuple float64
}

const maxStatSize = 20

// GpmonStats is a pack of multiple GpmonStat.
type GpmonStats struct {
	Length int64
	Data   [maxStatSize]GpmonStat
}

func (s *GpmonStat) ToPitMessage() (protos.PitMessage, error) {
	return protos.PitMessage{
		PitType: protos.EnumPitType_INSTRUMENT_INFO,
		Message: &protos.PitMessage_Instrument{
			Instrument: &protos.PlumInstrument{
				Running:    s.Running,
				TupleCount: s.TupleCount,
				Nloops:     s.NLoops,
				Ntuples:    s.NTuples,
				FirstTuple: s.FirstTuple,
				Key: &protos.DistributedNodeKey{
					SegId:  int32(s.SegID),
					ProcId: s.PID,
					NodeId: s.NID,
				},
			},
		},
	}, nil
}
