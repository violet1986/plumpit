package base

import (
	"plumpit/protos"
)

type RawMessage interface {
	ToPitMessage() (protos.PitMessage, error)
}

type Source interface{}
type SystemSource interface {
	GetSystemCpu() (protos.PitMessage, error)
}
type ProcSource interface {
	GetProcCpuPercent(duration interface{}) (protos.PitMessage, error)
	GetProcMemInfo() (protos.PitMessage, error)
	GetProcMemPercent() (protos.PitMessage, error)
}

type InstrumentSource interface {
	GetInstrumentInfo([]SharedMemoryOffset) (protos.PitMessage, error)
}

// ActiveSourceServer represent source that run standalone, receive query related message and then husk the message into PitMessage.
type ActiveSourceServer interface {
	Run(args ...interface{}) error
	GetRawMessage(Unmarshaller) (RawMessage, error)
}

type MapKey interface {
	GetHashKeyString() string
}
