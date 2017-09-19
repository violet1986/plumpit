package base

import (
	"plumpit/protos"
	"time"
)

type Unmarshaller func([]byte) (RawMessage, error)
type MultiUnmarshaller func([]byte) ([]RawMessage, error)

type Sender func(protos.PitMessage) error

type RuntimeConfig struct {
	MemSource            *SystemSource
	CPUSource            *SystemSource
	IOSource             *SystemSource
	CollatorEmitInterval time.Duration
	ProcIDs              map[int32]bool
}
