package querypit

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func StartUDPServer(addr string) (*net.UDPConn, error) {
	address, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return net.ListenUDP("udp", address)
}

func UDPTestClient(addr string) error {
	ServerAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		return err
	}
	defer Conn.Close()
	i := 0
	for i < 10 {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		fmt.Println(string(buf))
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}
