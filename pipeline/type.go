package pipeline

type PacketType int

type RawPacketData interface {
	getType() PacketType
	convertTo()
}

type PitMessage interface {
}
