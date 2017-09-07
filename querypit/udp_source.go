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
		msg, err := s.GetRawMessage(udpUnmarshallerForGpmonPkt)
		// error should return?
		if err != nil {
			return err
		}
		if msg != nil {
			msg.ToPitMessage()
		}
	}
}
func contentUnmarshallerForGpmonPkt(pkttype int, buf []byte) (base.RawMessage, error) {
	switch pkttype {
	case GpmonPktTypeQlog:
		pack := GpmonQlog{}
		err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &pack)
		return pack, err

	case GpmonPktTypeQexec:
		pack := GpmonQexec{}
		err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &pack)
		return pack, err
	default:
	}

	return nil, nil
}

func udpUnmarshallerForGpmonPkt(buf []byte) (base.RawMessage, error) {
	prefix := GpmonPacket{}
	err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &prefix)
	if err != nil {
		return nil, err
	}
	return contentUnmarshallerForGpmonPkt(int(prefix.Pkttype), buf[unsafe.Sizeof(prefix):len(buf)])
}
