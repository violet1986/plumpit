package base

// DistributedNodeKey represent a node index
type DistributedNodeKey struct {
	SegID  int
	PID    int
	NodeID int
}

// DistributedNode represent a node executed in segment/process
type DistributedNode struct {
	Key           DistributedNodeKey
	InstruMetrics RawPlumInstrument
}

// PlanNodeMetrics represents the general planned information of a plan node
type PlanNodeMetrics struct {
	PlanWidth   int     `json:"PlanWidth"`
	StartupCost float64 `json:"StartupCost"`
	TotalCost   float64 `json:"TotalCost"`
	PlanRows    float64 `json:"PlanRows"`
}

// PlumPlanNode represents a greenplum query plan node
type PlumPlanNode struct {
	SubNodes    []*DistributedNode `json:"Nodes"`
	Children    []*PlumPlanNode    `json:"Children"`
	PlanMetrics PlanNodeMetrics    `json:"PlanMetrics"`
}

// PlumQuery represent a greenplum query status
type PlumQuery struct {
	QueryID     string       `json:"id"`
	QueryStatus int          `json:"status"`
	QueryPlan   PlumPlanNode `json:"plan"`
}

const (
	GpmonPktTypeNone = iota
	GpmonPktTypeHello
	GpmonPktTypeMetrics
	GpmonPktTypeQlog
	GpmonPktTypeQexec
)
