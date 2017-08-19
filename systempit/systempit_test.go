package systempit_test

import (
	"plumpit/base"
	"plumpit/protos"
	. "plumpit/systempit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Systempit With Sigar Source", func() {
	var sigarTestSource SigarSource
	BeforeEach(func() {
		sigarTestSource = SigarSource{}
	})
	Context("System CPU test", func() {
		functions := []base.SourceHandlerFunc{base.GetSystemCPUHandler}
		res := GetSystemGoGenerators(functions)
		It("Test GetSystemGoGenerators", func() {
			Expect(len(res)).NotTo(Equal(0))
		})
		It("Test GetSystemCPU", func() {
			getcpu := res[0]
			ch := make(chan interface{})
			go getcpu(sigarTestSource, ch)
			cpu := <-ch
			syscpu := cpu.(protos.SystemCPU)
			//Expect(syscpu.Sys == 0).To(BeFalse())
			Expect(syscpu.PitType).To(Equal(protos.EnumPitType_SYSTEM_CPU))
		})

	})

})
