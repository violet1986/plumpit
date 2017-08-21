package systempit_test

import (
	"plumpit/base"
	"plumpit/protos"
	. "plumpit/systempit"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Systempit With Sigar Source", func() {
	var sigarTestSource SigarSource
	BeforeEach(func() {
		sigarTestSource = SigarSource{}
	})
	Context("System CPU test", func() {
		cpuFunc := base.GetSystemCpuDelegator()
		It("Test GetSystemCpuDelegator", func() {
			Expect(cpuFunc).NotTo(BeNil())
		})
		It("Test GetSystemCpu", func() {
			cpu, err := cpuFunc(sigarTestSource)
			Expect(err).NotTo(HaveOccurred())
			systemCpu := cpu.GetSystemCpu()
			Expect(systemCpu).NotTo(BeNil())
			Expect(systemCpu.Sys == 0).To(BeFalse())
			Expect(systemCpu.PitType).To(Equal(protos.EnumPitType_SYSTEM_CPU))
		})

	})

})

var _ = Describe("Proc with gopsutil", func() {
	//testPid :=
	testPid := 8084
	procTestSource := NewProcPsutilSource(testPid)
	getCpuEvery3Second := base.GetProcCpuDelegator(time.Duration(3))
	It("process should not be nil", func() {
		Expect(procTestSource).NotTo(BeNil())
	})
	It("process cpu should not be 0", func() {
		Expect(getCpuEvery3Second).NotTo(BeNil())
		cpu, err := getCpuEvery3Second(procTestSource)
		Expect(err).NotTo(HaveOccurred())
		Expect(cpu.GetProcCpu()).NotTo(BeNil())
		Expect(cpu.GetProcCpu().Percent == 0.0).NotTo(BeTrue())
	})

})
