package querypit

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"plumpit/base"
	"time"

	"github.com/ghetzel/shmtool/shm"
)

var gShmSourceMap = map[base.SharedMemoryInfo]*GPShmSource{}
var gRunningQueriesShmInfo = map[string]base.SharedMemoryInfo{}

const gpInstrSize = 168

type GPShmSource struct {
	ShmID             base.SharedMemoryInfo
	InstrumentOffsets map[string]int64
	quit              chan int
	// A lock?
}

func (s *GPShmSource) GetRawMessage(seg *shm.Segment, offset int64, size int64, unpacker base.Unmarshaller) (base.RawMessage, error) {
	buf, err := seg.ReadChunk(size, offset)
	if err != nil {
		return nil, err
	}
	return unpacker(buf)
}

func (s *GPShmSource) UnpackShm(seg *shm.Segment, size int64, unpacker base.Unmarshaller) func(offset map[string]int64) {
	return func(offsets map[string]int64) {
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
func (s *GPShmSource) Run(args ...interface{}) error {
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
	for {
		select {
		case <-ticker.C:
			instruHandler := s.UnpackShm(seg, gpInstrSize, func(buf []byte) (base.RawMessage, error) {
				pack := GPInstrument{}
				fmt.Println(hex.Dump(buf))
				err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &pack)
				return pack, err
			})
			instruHandler(s.InstrumentOffsets)
		case <-s.quit:
			ticker.Stop()
			return nil
		}
	}
}
