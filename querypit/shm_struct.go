package querypit

import (
	"fmt"
	"plumpit/protos"
)

type GPInstrument struct {
	Running    bool
	Dummy1     [48]byte
	TupleCount float64 `json:"TupleCount"`
	Dummy2     [8]byte
	NTuples    float64 `json:"NTuples"`
	NLoops     float64 `json:"NLoops"`
}

func (i GPInstrument) ToPitMessage() (protos.PitMessage, error) {
	fmt.Println(i.Running, i.TupleCount)
	return protos.PitMessage{}, nil
}
