package systempit_test

import (
	"fmt"
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
		cpuFunc := base.GetSystemCPUHandler()
		It("Test GetSystemGoGenerators", func() {
			Expect(cpuFunc).NotTo(BeNil())
		})
		It("Test GetSystemCPU", func() {
			cpu := cpuFunc(sigarTestSource)
			Expect(cpu).NotTo(BeNil())
			syscpu := cpu.(protos.SystemCPU)
			Expect(syscpu.Sys == 0).To(BeFalse())
			Expect(syscpu.PitType).To(Equal(protos.EnumPitType_SYSTEM_CPU))
		})

	})

})

var _ = Describe("Proc with gopsutil", func() {
	//testPid :=
	testPid := 8084
	procTestSource := NewProcPsutilSource(testPid)
	getCpuEvery3Second := base.GetProcCPUHandler(time.Duration(3))
	It("process should not be nil", func() {
		Expect(procTestSource).NotTo(BeNil())
	})
	It("process cpu should not be 0", func() {
		Expect(getCpuEvery3Second).NotTo(BeNil())
		cpu := getCpuEvery3Second(procTestSource).(protos.ProcCPU)
		Expect(cpu.Percent == 0.0).NotTo(BeTrue(), fmt.Sprintf("%f", cpu.Percent))
	})

})
