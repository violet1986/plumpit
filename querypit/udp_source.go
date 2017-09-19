package querypit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"plumpit/base"
	"plumpit/collator"
	"unsafe"
)

type UdpSource struct {
	conn *net.UDPConn
}

const messageSize = 1024

// GetRawMessage is not supported in UdpSource.
func (s UdpSource) GetRawMessage(unpacker base.Unmarshaller) (base.RawMessage, error) {
	return nil, nil
}

// GetRawMessages return an array of RawMessages.
func (s UdpSource) GetRawMessages(unpacker base.MultiUnmarshaller) ([]base.RawMessage, error) {
	buf := make([]byte, messageSize)
	n, _, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}
	return unpacker(buf[0:n])
}

// Run The run function for UdpSource
// args[0]: address
// args[1]: quit channel
// args[2]: Collator
func (s UdpSource) Run(args ...interface{}) error {
	var err error
	if len(args) < 3 {
		return fmt.Errorf("no enough argument to run udp source")
	}
	s.conn, err = StartUDPServer(args[0].(string))
	if err != nil {
		return err
	}
	defer s.conn.Close()
	log.Println("Start listening udp packets at", args[0].(string))
	quit := args[1].(chan int)
	msgCollator := args[2].(base.Collator).AddMessageFunc(collator.PrintSender)
	for {
		select {
		case command := <-quit:
			if command > 0 {
				log.Println("Quit command received, will quit now...")
				return nil
			}
		default:
			msgs, err := s.GetRawMessages(udpUnmarshallerForGpmonPkt)
			if err != nil {
				continue
			}
			if len(msgs) == 0 {
				continue
			}
			for _, m := range msgs {
				go msgCollator(m)
			}
		}
	}
}
func contentUnmarshallerForGpmonPkt(pkttype int, buf []byte) ([]base.RawMessage, error) {
	result := []base.RawMessage{}
	var pack base.RawMessage
	var err error
	switch pkttype {
	case gpmonPktTypeQlog:
		pack = &GpmonQlog{}
		err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, pack)
	case gpmonPktTypeQexec:
		pack = &GpmonQexec{}
		err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, pack)
	case gpmonPktTypeStat:
		stats := &GpmonStats{}
		err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, stats)
		var i int64
		for i = 0; i < stats.Length; i++ {
			result = append(result, &stats.Data[i])
		}
	default:
	}
	if err == nil && pack != nil {
		result = append(result, pack)
	}
	return result, err
}

func udpUnmarshallerForGpmonPkt(buf []byte) ([]base.RawMessage, error) {
	prefix := GpmonPacket{}
	err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &prefix)
	if err != nil {
		return nil, err
	}
	return contentUnmarshallerForGpmonPkt(int(prefix.Pkttype), buf[unsafe.Sizeof(prefix):len(buf)])
}
