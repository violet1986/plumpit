package querypit

import (
	"fmt"
	"plumpit/base"
	"time"

	"github.com/ghetzel/shmtool/shm"
)

type GPShmSource struct {
	ShmID   base.SharedMemoryInfo
	Offsets map[base.DistributedNodeKey]uint64
	// A lock?
}

func (s GPShmSource) GetRawMessage(seg *shm.Segment, offset uint64) (base.RawMessage, error) {
	seg.Seek(int64(offset), 0)
	instru := base.RawPlumInstrument{}
	seg.Read(instru)
	return nil, nil
}

// Run of GPShmSource:
// args[0] time interval to get shared memory.
// args[1] the quit channel
func (s GPShmSource) Run(args ...interface{}) error {
	seg, err := shm.Open(int(s.ShmID))
	if err != nil {
		return err
	}
	if seg == nil {
		return fmt.Errorf("Shm Open get nil")
	}
	p, err := seg.Attach()
	defer seg.Detach(p)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(args[0].(time.Duration) * time.Second)
	quit := args[1].(chan int)
	for {
		select {
		case <-ticker.C:
			for _, offset := range s.Offsets {
				instrumentMsg, err := s.GetRawMessage(seg, offset)
				if err != nil {
					continue
				}
				if instrumentMsg != nil {
					instrumentMsg.ToPitMessage()
				}
			}
		case <-quit:
			ticker.Stop()
			return nil
		}
	}
}
