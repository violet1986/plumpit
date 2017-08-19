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
	GetProcCPU() protos.ProcCPU
}
type SourceHandlerFunc func(s Source) interface{}

func GetSystemCPUHandler(s Source) interface{} {
	if sysSource, ok := s.(SystemSource); ok {
		return sysSource.GetSystemCPU()
	}
	return nil
}
