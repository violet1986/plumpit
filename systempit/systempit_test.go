package systempit_test

import (
	"fmt"
	"os"
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
		res := base.GoGenerators(functions)
		It("Test GetSystemGoGenerators", func() {
			Expect(len(res)).NotTo(Equal(0))
		})
		It("Test GetSystemCPU", func() {
			getcpu := res[0]
			ch := make(chan interface{})
			go getcpu(sigarTestSource, ch)
			cpu := <-ch
			Expect(cpu).NotTo(BeNil())
			syscpu := cpu.(protos.SystemCPU)
			Expect(syscpu.Sys == 0).To(BeFalse())
			Expect(syscpu.PitType).To(Equal(protos.EnumPitType_SYSTEM_CPU))
		})

	})

})

var _ = Describe("Proc with gopsutil", func() {
	testPid := os.Getpid()
	proc := NewProcPsutilSource(testPid)
	cpu := (*proc).GetProcCPU()
	fmt.Println(cpu)
	It("process should not be nil", func() {
		Expect(proc).NotTo(BeNil())
	})
	It("process cpu should not be 0", func() {
		Expect(cpu).NotTo(BeNil())
		Expect(cpu.Percent == 0.0).NotTo(BeTrue(), fmt.Sprintf("%f", cpu.Percent))
	})

})
