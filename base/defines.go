package base

import "plumpit/protos"

type RawType int

type RawPit interface {
	getType() RawType
	Husk()
}

type PitMessage interface {
	Append(interface{}) PitMessage
}
type Source interface{}
type SystemSource interface {
	GetSystemCPU() protos.SystemCPU
}
type ProcSource interface {
	GetProcCPU(duration interface{}) protos.ProcCPU
}
