package collator

import "plumpit/protos"

type RawCollator struct{}

func (r *RawCollator) AddMessage(msg protos.PitMessage) error {
	//Send it out directly
	return nil
}
