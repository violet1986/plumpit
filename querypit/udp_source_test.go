package querypit_test

import (
	. "plumpit/querypit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UDP source", func() {
	address := ":8888"
	udpSource := UdpSource{}
	err := udpSource.Run(address)
	It("could run", func() {
		Expect(err).NotTo(HaveOccurred())
		err = UDPTestClient(address)
		Expect(err).NotTo(HaveOccurred())
	})
})
