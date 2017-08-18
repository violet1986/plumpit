package main

import (
	"flag"
	"log"
	"net"
	"plumpit/gpmonpacket"

	"github.com/gogo/protobuf/proto"
)

var (
	mode = flag.String("m", "server", "mode: client or server")
	port = flag.String("p", "4000", "host: ip:port")
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func main() {
	RunServer()
}

func RunServer() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":8888")
	CheckError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()

	buf := make([]byte, 1024)

	log.Println("Listening on port " + *port)
	for {
		n, _, err := serverConn.ReadFromUDP(buf)
		packet := &gpmonpacket.GpmonPacket{}
		err = proto.Unmarshal(buf[0:n], packet)
		log.Printf("Packet magic is %d, type is %d", packet.Magic, packet.Pkttype)

		if err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
