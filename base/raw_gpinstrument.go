package base

type instrTime [16]byte
type SharedMemoryInfo uint64

// Instrument represents a Instrumentation struct in Greenplum
type RawPlumInstrument struct {
	Key           *DistributedNodeKey
	ShmemID       SharedMemoryInfo
	Running       bool    `json:"Running"`
	FirstTuple    float64 `json:"FirstTuple"`
	TupleCount    float64 `json:"TupleCount"`
	Startup       float64 `json:"Startup"`
	Total         float64 `json:"Total"`
	NTuples       float64 `json:"NTuples"`
	NLoops        float64 `json:"NLoops"`
	ExecMemUsed   float64 `json:"-"`
	WorkMemUsed   float64 `json:"-"`
	WorkMemWanted float64 `json:"-"`
}
