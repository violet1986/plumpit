package querypit

import (
	"bytes"
	"encoding/binary"
	"net"
	"plumpit/base"
	"unsafe"
)

type UdpSource struct {
	conn *net.UDPConn
}

const messageSize = 1024

func (s UdpSource) GetRawMessage(unpacker base.Unmarshaller) (base.RawMessage, error) {
	buf := make([]byte, messageSize)
	n, _, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}
	return unpacker(buf[0:n])
}

// Run The run function for UdpSource
// args[0]: address
func (s UdpSource) Run(args ...interface{}) error {
	var err error
	s.conn, err = StartUDPServer(args[0].(string))
	if err != nil {
		return err
	}
	defer s.conn.Close()
	for {
		_, err := s.GetRawMessage(udpUnmarshallerForGpmonPkt)
		if err != nil {
			return err
		}
		//msg.ToPitMessage()
	}
}
func getSubUnmarshallerForGpmonPkt(pkttype int, args ...interface{}) base.Unmarshaller {
	switch pkttype {
	default:
		return func([]byte) (base.RawMessage, error) {
			return nil, nil
		}
	}
}

func udpUnmarshallerForGpmonPkt(buf []byte) (base.RawMessage, error) {
	prefix := GpmonPacket{}
	err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &prefix)
	if err != nil {
		return nil, err
	}
	realPacketUnpaker := getSubUnmarshallerForGpmonPkt(int(prefix.Pkttype))
	return realPacketUnpaker(buf[unsafe.Sizeof(prefix):len(buf)])
}
