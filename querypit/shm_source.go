package querypit

import (
	"fmt"
	"plumpit/base"
	"time"

	"github.com/ghetzel/shmtool/shm"
)

type GPShmSource struct {
	ShmID             base.SharedMemoryInfo
	InstrumentOffsets map[base.DistributedNodeKey]uint64
	// A lock?
}

func (s GPShmSource) GetRawMessage(seg *shm.Segment, offset uint64, size uint64, unpacker base.Unmarshaller) (base.RawMessage, error) {
	seg.Seek(int64(offset), 0)
	buf := [size]byte{}
	seg.Read(buf)
	return unpacker(buf)
}

func UnpackShm(seg *shm.Segment, offsets map[interface{}]uint64) func(size uint64, base.Unmarshaller) {
	return func(size uint64, unpacker base.Unmarshaller) {
		for _, offset := range offsets {
			msg, err := s.GetRawMessage(seg, offset, size, unpacker)
			if err != nil {
				continue
			}
			if msg != nil {
				msg.ToPitMessage()
			}
		}
	}
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
			instruHandler := UnpackShm(seg, s.InstrumentOffsets)
			instruHandler(unsafe.Sizeof(base.RawPlumInstrument), func(b []byte) (RawMessage, error) {
				
			})
		case <-quit:
			ticker.Stop()
			return nil
		}
	}
}
