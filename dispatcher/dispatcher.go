package dispatcher

/*
dispatch is the scheduled runner of system metrics collection and control when should collator emit itself
*/
import (
	"plumpit/base"
	"time"
)

func NewDispatcherRunner(c base.Collator, sender base.Sender) func(base.RuntimeConfig) {
	return func(conf base.RuntimeConfig) {
		collateTicker := time.NewTicker(conf.CollatorEmitInterval)
		memFunc := base.GetProcMemPercentDelegator()
		for {
			select {
			case <-collateTicker.C:
				mem, err := memFunc(conf.MemSource)
				if err == nil {
					c.Collate(mem)
				}
				pit, err := c.ToPitMessage()
				if err == nil {
					sender(pit)
				}
			}
		}
	}
}
