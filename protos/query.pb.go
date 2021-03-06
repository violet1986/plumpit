// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EnumQueryStatus int32

const (
	EnumQueryStatus_QUERY_SILENT    EnumQueryStatus = 0
	EnumQueryStatus_QUERY_SUBMIT    EnumQueryStatus = 1
	EnumQueryStatus_QUERY_START     EnumQueryStatus = 2
	EnumQueryStatus_QUERY_DONE      EnumQueryStatus = 3
	EnumQueryStatus_QUERY_ERROR     EnumQueryStatus = 4
	EnumQueryStatus_QUERY_CANCELING EnumQueryStatus = 5
)

var EnumQueryStatus_name = map[int32]string{
	0: "QUERY_SILENT",
	1: "QUERY_SUBMIT",
	2: "QUERY_START",
	3: "QUERY_DONE",
	4: "QUERY_ERROR",
	5: "QUERY_CANCELING",
}
var EnumQueryStatus_value = map[string]int32{
	"QUERY_SILENT":    0,
	"QUERY_SUBMIT":    1,
	"QUERY_START":     2,
	"QUERY_DONE":      3,
	"QUERY_ERROR":     4,
	"QUERY_CANCELING": 5,
}

func (x EnumQueryStatus) String() string {
	return proto.EnumName(EnumQueryStatus_name, int32(x))
}
func (EnumQueryStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type EnumNodeStatus int32

const (
	EnumNodeStatus_NODE_INIT  EnumNodeStatus = 0
	EnumNodeStatus_NODE_START EnumNodeStatus = 1
	EnumNodeStatus_NODE_DONE  EnumNodeStatus = 2
)

var EnumNodeStatus_name = map[int32]string{
	0: "NODE_INIT",
	1: "NODE_START",
	2: "NODE_DONE",
}
var EnumNodeStatus_value = map[string]int32{
	"NODE_INIT":  0,
	"NODE_START": 1,
	"NODE_DONE":  2,
}

func (x EnumNodeStatus) String() string {
	return proto.EnumName(EnumNodeStatus_name, int32(x))
}
func (EnumNodeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type DistributedNodeKey struct {
	SegId  int32 `protobuf:"varint,1,opt,name=seg_id,json=segId" json:"seg_id,omitempty"`
	ProcId int32 `protobuf:"varint,2,opt,name=proc_id,json=procId" json:"proc_id,omitempty"`
	NodeId int32 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *DistributedNodeKey) Reset()                    { *m = DistributedNodeKey{} }
func (m *DistributedNodeKey) String() string            { return proto.CompactTextString(m) }
func (*DistributedNodeKey) ProtoMessage()               {}
func (*DistributedNodeKey) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *DistributedNodeKey) GetSegId() int32 {
	if m != nil {
		return m.SegId
	}
	return 0
}

func (m *DistributedNodeKey) GetProcId() int32 {
	if m != nil {
		return m.ProcId
	}
	return 0
}

func (m *DistributedNodeKey) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

type PlumInstrument struct {
	Key        *DistributedNodeKey `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Running    bool                `protobuf:"varint,2,opt,name=running" json:"running,omitempty"`
	TupleCount uint64              `protobuf:"varint,3,opt,name=tuple_count,json=tupleCount" json:"tuple_count,omitempty"`
	Ntuples    uint64              `protobuf:"varint,4,opt,name=ntuples" json:"ntuples,omitempty"`
	Nloops     uint64              `protobuf:"varint,5,opt,name=nloops" json:"nloops,omitempty"`
	FirstTuple float64             `protobuf:"fixed64,6,opt,name=first_tuple,json=firstTuple" json:"first_tuple,omitempty"`
}

func (m *PlumInstrument) Reset()                    { *m = PlumInstrument{} }
func (m *PlumInstrument) String() string            { return proto.CompactTextString(m) }
func (*PlumInstrument) ProtoMessage()               {}
func (*PlumInstrument) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PlumInstrument) GetKey() *DistributedNodeKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PlumInstrument) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *PlumInstrument) GetTupleCount() uint64 {
	if m != nil {
		return m.TupleCount
	}
	return 0
}

func (m *PlumInstrument) GetNtuples() uint64 {
	if m != nil {
		return m.Ntuples
	}
	return 0
}

func (m *PlumInstrument) GetNloops() uint64 {
	if m != nil {
		return m.Nloops
	}
	return 0
}

func (m *PlumInstrument) GetFirstTuple() float64 {
	if m != nil {
		return m.FirstTuple
	}
	return 0
}

type QueryInfo struct {
	QueryId    string                     `protobuf:"bytes,1,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	Database   string                     `protobuf:"bytes,2,opt,name=database" json:"database,omitempty"`
	User       string                     `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	SubmitTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	StartTime  *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	FinishTime *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	Status     EnumQueryStatus            `protobuf:"varint,7,opt,name=status,enum=protos.EnumQueryStatus" json:"status,omitempty"`
}

func (m *QueryInfo) Reset()                    { *m = QueryInfo{} }
func (m *QueryInfo) String() string            { return proto.CompactTextString(m) }
func (*QueryInfo) ProtoMessage()               {}
func (*QueryInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *QueryInfo) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *QueryInfo) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *QueryInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryInfo) GetSubmitTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.SubmitTime
	}
	return nil
}

func (m *QueryInfo) GetStartTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *QueryInfo) GetFinishTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

func (m *QueryInfo) GetStatus() EnumQueryStatus {
	if m != nil {
		return m.Status
	}
	return EnumQueryStatus_QUERY_SILENT
}

type ExecInfo struct {
	QueryId     string              `protobuf:"bytes,1,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	NodeKey     *DistributedNodeKey `protobuf:"bytes,2,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	Status      EnumNodeStatus      `protobuf:"varint,3,opt,name=status,enum=protos.EnumNodeStatus" json:"status,omitempty"`
	StartupCost float64             `protobuf:"fixed64,4,opt,name=startup_cost,json=startupCost" json:"startup_cost,omitempty"`
	TotalCost   float64             `protobuf:"fixed64,5,opt,name=total_cost,json=totalCost" json:"total_cost,omitempty"`
	PlanRows    float64             `protobuf:"fixed64,6,opt,name=plan_rows,json=planRows" json:"plan_rows,omitempty"`
	NodeType    int32               `protobuf:"varint,7,opt,name=node_type,json=nodeType" json:"node_type,omitempty"`
}

func (m *ExecInfo) Reset()                    { *m = ExecInfo{} }
func (m *ExecInfo) String() string            { return proto.CompactTextString(m) }
func (*ExecInfo) ProtoMessage()               {}
func (*ExecInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ExecInfo) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *ExecInfo) GetNodeKey() *DistributedNodeKey {
	if m != nil {
		return m.NodeKey
	}
	return nil
}

func (m *ExecInfo) GetStatus() EnumNodeStatus {
	if m != nil {
		return m.Status
	}
	return EnumNodeStatus_NODE_INIT
}

func (m *ExecInfo) GetStartupCost() float64 {
	if m != nil {
		return m.StartupCost
	}
	return 0
}

func (m *ExecInfo) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ExecInfo) GetPlanRows() float64 {
	if m != nil {
		return m.PlanRows
	}
	return 0
}

func (m *ExecInfo) GetNodeType() int32 {
	if m != nil {
		return m.NodeType
	}
	return 0
}

func init() {
	proto.RegisterType((*DistributedNodeKey)(nil), "protos.DistributedNodeKey")
	proto.RegisterType((*PlumInstrument)(nil), "protos.PlumInstrument")
	proto.RegisterType((*QueryInfo)(nil), "protos.QueryInfo")
	proto.RegisterType((*ExecInfo)(nil), "protos.ExecInfo")
	proto.RegisterEnum("protos.EnumQueryStatus", EnumQueryStatus_name, EnumQueryStatus_value)
	proto.RegisterEnum("protos.EnumNodeStatus", EnumNodeStatus_name, EnumNodeStatus_value)
}

func init() { proto.RegisterFile("query.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x9c, 0xc4, 0x93, 0x92, 0x5a, 0x8b, 0x68, 0x43, 0x11, 0x6a, 0xc9, 0xa9, 0xaa,
	0x50, 0x2a, 0x15, 0x71, 0x40, 0x95, 0x90, 0x4a, 0x6a, 0x21, 0x8b, 0xe2, 0xd2, 0xad, 0x7b, 0xe0,
	0x80, 0x2c, 0x27, 0xde, 0x06, 0x8b, 0x78, 0xd7, 0x78, 0x77, 0x55, 0x22, 0xf1, 0x01, 0x7c, 0x14,
	0x9f, 0xc0, 0x47, 0xa1, 0x9d, 0xb5, 0x4b, 0x0b, 0x87, 0x9c, 0xda, 0x79, 0xef, 0x8d, 0xe7, 0xed,
	0x9b, 0x51, 0x60, 0xf0, 0x4d, 0xb3, 0x6a, 0x35, 0x29, 0x2b, 0xa1, 0x04, 0xe9, 0xe2, 0x1f, 0xb9,
	0xb3, 0xbb, 0x10, 0x62, 0xb1, 0x64, 0x87, 0x58, 0xce, 0xf4, 0xf5, 0xa1, 0xca, 0x0b, 0x26, 0x55,
	0x5a, 0x94, 0x56, 0x38, 0xfe, 0x0c, 0xe4, 0x34, 0x97, 0xaa, 0xca, 0x67, 0x5a, 0xb1, 0x2c, 0x12,
	0x19, 0x7b, 0xcf, 0x56, 0xe4, 0x31, 0x74, 0x25, 0x5b, 0x24, 0x79, 0x36, 0x72, 0xf6, 0x9c, 0x7d,
	0x97, 0xba, 0x92, 0x2d, 0xc2, 0x8c, 0x6c, 0x43, 0xaf, 0xac, 0xc4, 0xdc, 0xe0, 0x2d, 0xc4, 0xcd,
	0x98, 0xb9, 0x25, 0xb8, 0xc8, 0x98, 0x21, 0xda, 0x96, 0x30, 0x65, 0x98, 0x8d, 0x7f, 0x3b, 0x30,
	0xfc, 0xb8, 0xd4, 0x45, 0xc8, 0xa5, 0xaa, 0x74, 0xc1, 0xb8, 0x22, 0x2f, 0xa0, 0xfd, 0x95, 0xad,
	0xf0, 0xc3, 0x83, 0xa3, 0x1d, 0x6b, 0x43, 0x4e, 0xfe, 0x37, 0x41, 0x8d, 0x8c, 0x8c, 0xa0, 0x57,
	0x69, 0xce, 0x73, 0xbe, 0xc0, 0x91, 0x7d, 0xda, 0x94, 0x64, 0x17, 0x06, 0x4a, 0x97, 0x4b, 0x96,
	0xcc, 0x85, 0xe6, 0x0a, 0xe7, 0x76, 0x28, 0x20, 0x34, 0x35, 0x88, 0x69, 0xe5, 0x58, 0xca, 0x51,
	0x07, 0xc9, 0xa6, 0x24, 0x5b, 0xd0, 0xe5, 0x4b, 0x21, 0x4a, 0x39, 0x72, 0x91, 0xa8, 0x2b, 0xf3,
	0xc9, 0xeb, 0xbc, 0x92, 0x2a, 0x41, 0xdd, 0xa8, 0xbb, 0xe7, 0xec, 0x3b, 0x14, 0x10, 0x8a, 0x0d,
	0x32, 0xfe, 0xd5, 0x02, 0xef, 0xc2, 0xc4, 0x1c, 0xf2, 0x6b, 0x41, 0x9e, 0x40, 0x1f, 0x33, 0x6f,
	0x72, 0xf2, 0x68, 0x0f, 0xeb, 0x30, 0x23, 0x3b, 0xd0, 0xcf, 0x52, 0x95, 0xce, 0x52, 0xc9, 0xd0,
	0xb7, 0x47, 0x6f, 0x6b, 0x42, 0xa0, 0xa3, 0x25, 0xab, 0xd0, 0xb1, 0x47, 0xf1, 0x7f, 0x72, 0x0c,
	0x03, 0xa9, 0x67, 0x45, 0xae, 0x12, 0xb3, 0x20, 0xf4, 0x6b, 0xc2, 0xb1, 0xdb, 0x9b, 0x34, 0xdb,
	0x9b, 0xc4, 0xcd, 0xf6, 0x28, 0x58, 0xb9, 0x01, 0xc8, 0x6b, 0x00, 0xa9, 0xd2, 0xaa, 0xee, 0x75,
	0xd7, 0xf6, 0x7a, 0xa8, 0xc6, 0xd6, 0x63, 0xf3, 0x62, 0x9e, 0xcb, 0x2f, 0xb6, 0xb7, 0xbb, 0x7e,
	0xae, 0x95, 0x63, 0xf3, 0x21, 0x74, 0xa5, 0x4a, 0x95, 0x96, 0xa3, 0xde, 0x9e, 0xb3, 0x3f, 0x3c,
	0xda, 0x6e, 0x96, 0x19, 0x70, 0x5d, 0x60, 0x4c, 0x97, 0x48, 0xd3, 0x5a, 0x36, 0xfe, 0xd9, 0x82,
	0x7e, 0xf0, 0x9d, 0xcd, 0xd7, 0xa5, 0xf7, 0x0a, 0xfa, 0x78, 0x4e, 0xe6, 0x4e, 0x5a, 0x6b, 0xef,
	0x04, 0x4f, 0xcf, 0x5c, 0xed, 0xe4, 0xd6, 0x4f, 0x1b, 0xfd, 0x6c, 0xdd, 0xf5, 0x63, 0xd4, 0xf7,
	0xed, 0x90, 0xe7, 0xb0, 0x81, 0x49, 0xe8, 0x32, 0x99, 0x0b, 0xa9, 0x30, 0x75, 0x87, 0x0e, 0x6a,
	0x6c, 0x2a, 0xa4, 0x22, 0xcf, 0x00, 0x94, 0x50, 0xe9, 0xd2, 0x0a, 0x5c, 0x14, 0x78, 0x88, 0x20,
	0xfd, 0x14, 0xbc, 0x72, 0x99, 0xf2, 0xa4, 0x12, 0x37, 0xb2, 0x3e, 0x97, 0xbe, 0x01, 0xa8, 0xb8,
	0x91, 0x86, 0xc4, 0x57, 0xa8, 0x55, 0xc9, 0x30, 0x21, 0x97, 0xe2, 0xb3, 0xe2, 0x55, 0xc9, 0x0e,
	0x7e, 0xc0, 0xe6, 0x3f, 0x29, 0x11, 0x1f, 0x36, 0x2e, 0xae, 0x02, 0xfa, 0x29, 0xb9, 0x0c, 0xcf,
	0x82, 0x28, 0xf6, 0x1f, 0xdc, 0x41, 0xae, 0xde, 0x7e, 0x08, 0x63, 0xdf, 0x21, 0x9b, 0x30, 0xa8,
	0x91, 0xf8, 0x84, 0xc6, 0x7e, 0x8b, 0x0c, 0x01, 0x2c, 0x70, 0x7a, 0x1e, 0x05, 0x7e, 0xfb, 0xaf,
	0x20, 0xa0, 0xf4, 0x9c, 0xfa, 0x1d, 0xf2, 0x08, 0x36, 0x2d, 0x30, 0x3d, 0x89, 0xa6, 0xc1, 0x59,
	0x18, 0xbd, 0xf3, 0xdd, 0x83, 0x37, 0x30, 0xbc, 0x9f, 0x09, 0x79, 0x08, 0x5e, 0x74, 0x7e, 0x1a,
	0x24, 0x61, 0x14, 0x9a, 0xc9, 0x43, 0x00, 0x2c, 0xed, 0x18, 0xe7, 0x96, 0xc6, 0x29, 0xad, 0x99,
	0xfd, 0x79, 0x79, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x49, 0x0b, 0x05, 0x85, 0x74, 0x04, 0x00,
	0x00,
}
