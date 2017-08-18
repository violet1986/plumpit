package plumquery

type instrTime [16]byte

// Instrument represents a Instrumentation struct in Greenplum
type Instrument struct {
	Key           *DistributedNodeKey
	ShmemID       int
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
