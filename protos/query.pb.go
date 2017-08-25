// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributedNodeKey struct {
	SegId  int32 `protobuf:"varint,1,opt,name=seg_id,json=segId" json:"seg_id,omitempty"`
	ProcId int32 `protobuf:"varint,2,opt,name=proc_id,json=procId" json:"proc_id,omitempty"`
	NodeId int32 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *DistributedNodeKey) Reset()                    { *m = DistributedNodeKey{} }
func (m *DistributedNodeKey) String() string            { return proto.CompactTextString(m) }
func (*DistributedNodeKey) ProtoMessage()               {}
func (*DistributedNodeKey) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

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
	Key           *DistributedNodeKey `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Running       bool                `protobuf:"varint,2,opt,name=running" json:"running,omitempty"`
	FirstTuple    float64             `protobuf:"fixed64,3,opt,name=first_tuple,json=firstTuple" json:"first_tuple,omitempty"`
	TupleCount    float64             `protobuf:"fixed64,4,opt,name=tuple_count,json=tupleCount" json:"tuple_count,omitempty"`
	StartUp       float64             `protobuf:"fixed64,5,opt,name=start_up,json=startUp" json:"start_up,omitempty"`
	Total         float64             `protobuf:"fixed64,6,opt,name=total" json:"total,omitempty"`
	Ntuples       float64             `protobuf:"fixed64,7,opt,name=ntuples" json:"ntuples,omitempty"`
	Nloops        float64             `protobuf:"fixed64,8,opt,name=nloops" json:"nloops,omitempty"`
	ExecMemUsed   float64             `protobuf:"fixed64,9,opt,name=exec_mem_used,json=execMemUsed" json:"exec_mem_used,omitempty"`
	WorkMemUsed   float64             `protobuf:"fixed64,10,opt,name=work_mem_used,json=workMemUsed" json:"work_mem_used,omitempty"`
	WorkMemWanted float64             `protobuf:"fixed64,11,opt,name=work_mem_wanted,json=workMemWanted" json:"work_mem_wanted,omitempty"`
}

func (m *PlumInstrument) Reset()                    { *m = PlumInstrument{} }
func (m *PlumInstrument) String() string            { return proto.CompactTextString(m) }
func (*PlumInstrument) ProtoMessage()               {}
func (*PlumInstrument) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

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

func (m *PlumInstrument) GetFirstTuple() float64 {
	if m != nil {
		return m.FirstTuple
	}
	return 0
}

func (m *PlumInstrument) GetTupleCount() float64 {
	if m != nil {
		return m.TupleCount
	}
	return 0
}

func (m *PlumInstrument) GetStartUp() float64 {
	if m != nil {
		return m.StartUp
	}
	return 0
}

func (m *PlumInstrument) GetTotal() float64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PlumInstrument) GetNtuples() float64 {
	if m != nil {
		return m.Ntuples
	}
	return 0
}

func (m *PlumInstrument) GetNloops() float64 {
	if m != nil {
		return m.Nloops
	}
	return 0
}

func (m *PlumInstrument) GetExecMemUsed() float64 {
	if m != nil {
		return m.ExecMemUsed
	}
	return 0
}

func (m *PlumInstrument) GetWorkMemUsed() float64 {
	if m != nil {
		return m.WorkMemUsed
	}
	return 0
}

func (m *PlumInstrument) GetWorkMemWanted() float64 {
	if m != nil {
		return m.WorkMemWanted
	}
	return 0
}

func init() {
	proto.RegisterType((*DistributedNodeKey)(nil), "protos.DistributedNodeKey")
	proto.RegisterType((*PlumInstrument)(nil), "protos.PlumInstrument")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x69, 0x6b, 0x92, 0x3a, 0xa1, 0x0a, 0x8b, 0x7f, 0x56, 0x2f, 0x4a, 0x0f, 0xe2, 0x41,
	0x7a, 0xd0, 0x47, 0xd0, 0x4b, 0x11, 0x45, 0x82, 0xc5, 0x93, 0x84, 0x36, 0x3b, 0x96, 0xd0, 0x64,
	0x37, 0xee, 0xce, 0x52, 0xfb, 0xd0, 0xbe, 0x83, 0xec, 0x6c, 0xab, 0x07, 0x4f, 0xe1, 0xfb, 0x7d,
	0xbf, 0xcc, 0x0c, 0x2c, 0xe4, 0x9f, 0x1e, 0xed, 0x66, 0xd2, 0x59, 0x43, 0x46, 0xa4, 0xfc, 0x71,
	0xe3, 0x77, 0x10, 0x0f, 0xb5, 0x23, 0x5b, 0x2f, 0x3c, 0xa1, 0x7a, 0x36, 0x0a, 0x1f, 0x71, 0x23,
	0x8e, 0x21, 0x75, 0xb8, 0x2c, 0x6b, 0x25, 0x7b, 0x97, 0xbd, 0xeb, 0xa4, 0x48, 0x1c, 0x2e, 0xa7,
	0x4a, 0x9c, 0x42, 0xd6, 0x59, 0x53, 0x05, 0xde, 0x67, 0x1e, 0xa6, 0x54, 0xb1, 0xd0, 0x46, 0x61,
	0x28, 0x06, 0xb1, 0x08, 0x71, 0xaa, 0xc6, 0xdf, 0x7d, 0x38, 0x78, 0x69, 0x7c, 0x3b, 0xd5, 0x8e,
	0xac, 0x6f, 0x51, 0x93, 0xb8, 0x81, 0xc1, 0x0a, 0x37, 0x3c, 0x38, 0xbf, 0x3d, 0x8f, 0xe7, 0xb8,
	0xc9, 0xff, 0x23, 0x8a, 0xa0, 0x09, 0x09, 0x99, 0xf5, 0x5a, 0xd7, 0x7a, 0xc9, 0x2b, 0x87, 0xc5,
	0x2e, 0x8a, 0x0b, 0xc8, 0x3f, 0x6a, 0xeb, 0xa8, 0x24, 0xdf, 0x35, 0xc8, 0x7b, 0x7b, 0x05, 0x30,
	0x7a, 0x0d, 0x24, 0x08, 0x5c, 0x95, 0x95, 0xf1, 0x9a, 0xe4, 0x5e, 0x14, 0x18, 0xdd, 0x07, 0x22,
	0xce, 0x60, 0xe8, 0x68, 0x6e, 0xa9, 0xf4, 0x9d, 0x4c, 0xb8, 0xcd, 0x38, 0xcf, 0x3a, 0x71, 0x04,
	0x09, 0x19, 0x9a, 0x37, 0x32, 0x65, 0x1e, 0x43, 0x38, 0x46, 0xf3, 0xff, 0x4e, 0x66, 0xd1, 0xdf,
	0x46, 0x71, 0x02, 0xa9, 0x6e, 0x8c, 0xe9, 0x9c, 0x1c, 0x72, 0xb1, 0x4d, 0x62, 0x0c, 0x23, 0xfc,
	0xc2, 0xaa, 0x6c, 0xb1, 0x2d, 0xbd, 0x43, 0x25, 0xf7, 0xb9, 0xce, 0x03, 0x7c, 0xc2, 0x76, 0xe6,
	0x50, 0x05, 0x67, 0x6d, 0xec, 0xea, 0xcf, 0x81, 0xe8, 0x04, 0xb8, 0x73, 0xae, 0xe0, 0xf0, 0xd7,
	0x59, 0xcf, 0x35, 0xa1, 0x92, 0x39, 0x5b, 0xa3, 0xad, 0xf5, 0xc6, 0x70, 0x11, 0x9f, 0xf5, 0xee,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x86, 0x18, 0x71, 0xbf, 0xec, 0x01, 0x00, 0x00,
}