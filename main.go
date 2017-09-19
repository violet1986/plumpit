package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"plumpit/collator"
	"plumpit/querypit"
	"syscall"
	"time"
)

var flagUDPPort int

func init() {
	flag.IntVar(&flagUDPPort, "udpport", 8888, "the udp server port")
}
func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	udpsource := querypit.UdpSource{}
	quit := make(chan int, 1)
	var err error
	coll := collator.NewProcessCollator()
	go func() {
		err = udpsource.Run(fmt.Sprintf(":%d", flagUDPPort), quit, coll)
	}()
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C:
				log.Println("coll is", coll.DebugString())
			}
		}
	}()
	for {
		s := <-signalChan
		switch s {
		case syscall.SIGHUP:
			fmt.Println(s)
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println(s)
			quit <- 1
			os.Exit(0)
		default:
			fmt.Println(s)
		}
	}
}
