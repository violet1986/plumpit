package base

import "plumpit/protos"

type RawType int

type RawPit interface {
	getType() RawType
	Husk()
}

type Source interface{}
type SystemSource interface {
	GetSystemCpu() protos.PitMessage
}
type ProcSource interface {
	GetProcCpu(duration interface{}) protos.PitMessage
	GetProcMem() protos.PitMessage
}
