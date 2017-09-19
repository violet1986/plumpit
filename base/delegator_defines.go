package base

import (
	"plumpit/protos"
	"time"
)

type Unmarshaller func([]byte) (RawMessage, error)
type MultiUnmarshaller func([]byte) ([]RawMessage, error)

type Sender func(protos.PitMessage) error

type RuntimeConfig struct {
	PSource              ProcSource
	CollatorEmitInterval time.Duration
	ProcIDs              map[int32]bool
}
