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
type QlogKey struct {
	Tmid int32
	Ssid int32
	Ccnt int32
}

type GpmonProcMetrics struct {
	FdCnt                    uint32
	CPUPct                   float32
	MemSize, Resident, Share uint64
}
type GpmonQlog struct {
	Key                   QlogKey
	User                  NameData
	Database              NameData
	Tsubmit, Tstart, Tfin int32
	Status                int32
	Cost                  int32
	CPUElapsed            int64
	PMetrics              GpmonProcMetrics
	SharedMemory          base.SharedMemoryInfo
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
