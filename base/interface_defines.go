package base

import (
	"plumpit/protos"
	"time"
)

type RawMessage interface {
	ToPitMessage() (protos.PitMessage, error)
}

type Source interface{}
type SystemSource interface {
	GetSystemCpu() (protos.PitMessage, error)
}
type ProcSource interface {
	GetProcCpuPercent(int32, time.Duration) (protos.PitMessage, error)
	GetProcMemInfo(int32) (protos.PitMessage, error)
	GetProcMemPercent(int32) (protos.PitMessage, error)
}

// ActiveSourceServer represent source that run standalone, receive query related message and then husk the message into PitMessage.
type ActiveSourceServer interface {
	Run(args ...interface{}) error
	GetRawMessage(Unmarshaller) (RawMessage, error)
	GetRawMessages(Unmarshaller) ([]RawMessage, error)
}

type MapKey interface {
	GetHashKeyString() string
}

// Collator can save RawMessage and transfer to PitMessage.
type Collator interface {
	AddMessageFunc(Sender) func(RawMessage) error
	ToPitMessage() (protos.PitMessage, error)
	Collate(protos.PitMessage)
}
