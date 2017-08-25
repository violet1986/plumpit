package querypit

import (
	"bytes"
	"encoding/binary"
	"net"
	"plumpit/base"
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

func udpUnmarshallerForGpmonPkt(buf []byte) (base.RawMessage, error) {
	p := GpmonPacket{}
	err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	switch p.Pkttype {

	}
	return nil, nil
}
