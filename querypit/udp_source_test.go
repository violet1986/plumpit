package querypit_test

import (
	"log"
	"plumpit/collator"
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
	coll := collator.NewProcessCollator()
	BeforeEach(func() {
		go func() {
			err = udpSource.Run(address, quit, coll)
		}()
	})
	It("could run", func() {
		ticker := time.NewTicker(1000 * time.Second)
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
