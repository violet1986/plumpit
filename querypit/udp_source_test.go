package querypit_test

import (
	"log"
	. "plumpit/querypit"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UDP source", func() {
	address := ":8888"
	udpSource := UdpSource{}
	quit := make(chan int, 1)
	var err error
	BeforeEach(func() {
		go func() {
			err = udpSource.Run(address, quit)
			log.Println("Start run")
		}()
	})
	It("could run", func() {
		ticker := time.NewTicker(5 * time.Second)
		select {
		case <-ticker.C:
			log.Println("Time up")
			quit <- 1
		}
		Expect(err).NotTo(HaveOccurred())
		//err = UDPTestClient(address)
		//Expect(err).NotTo(HaveOccurred())
	})
})
