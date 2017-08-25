package querypit

import (
	"plumpit/base"
)

type GpmonPacket struct {
	Magic         int32
	Version       int16
	Pkttype       int16
	ShmDescriptor base.SharedMemoryInfo
}
