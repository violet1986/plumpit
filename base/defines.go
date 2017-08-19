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

type Source interface {
	GetSystemCPU() protos.SystemCPU
}
type SourceHandlerFunc func(s Source) interface{}

func GetSystemCPUHandler(s Source) interface{} {
	return s.GetSystemCPU()
}
