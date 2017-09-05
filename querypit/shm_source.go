package querypit

import (
	"plumpit/base"
)

type GPShmSource struct {
	ShmID base.SharedMemoryInfo
}

func (s GPShmSource) GetRawMessage() (base.RawMessage, error) {
	return nil, nil
}
