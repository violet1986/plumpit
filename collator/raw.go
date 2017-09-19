package collator

import "plumpit/base"

type RawCollator struct{}

func (r *RawCollator) AddMessageFunc(sender base.Sender) func(base.RawMessage) error {
	return func(msg base.RawMessage) error {
		if msg != nil {
			pit, err := msg.ToPitMessage()
			if err != nil {
				return err
			}
			if r != nil {
				return sender(pit)
			}
		}
		return nil
	}
}
