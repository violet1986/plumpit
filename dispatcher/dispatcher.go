package dispatcher

/*
dispatch is the scheduled runner of system metrics collection and control when should collator emit itself
*/
import (
	"plumpit/base"
	"time"
)

func NewDispatcherRunner(conf *base.RuntimeConfig, c base.Collator) func(base.Sender) {
	collateTicker := time.NewTicker(conf.CollatorEmitInterval)
	GenerateProcMetrics := func() {
		for pid, ok := range conf.ProcIDs {
			if !ok {
				continue
			}
			mem, err := conf.PSource.GetProcMemInfo(pid)
			if err == nil {
				c.Collate(mem)
			}
		}
	}

	return func(s base.Sender) {
		for {
			select {
			case <-collateTicker.C:
				GenerateProcMetrics()
				pit, err := c.ToPitMessage()
				if err == nil {
					s(pit)
				}
			}
		}
	}
}
