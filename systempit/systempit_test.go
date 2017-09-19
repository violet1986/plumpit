package systempit_test

import (
	"fmt"
	"os"
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
		It("Test GetSystemCpu", func() {
			cpu, err := sigarTestSource.GetSystemCpu()
			Expect(err).NotTo(HaveOccurred())
			Expect(cpu.PitType).To(Equal(protos.EnumPitType_SYSTEM_CPU))
			systemCpu := cpu.GetSystemCpu()
			Expect(systemCpu).NotTo(BeNil())
			Expect(systemCpu.Sys == 0).To(BeFalse())
		})

	})

})

var _ = Describe("Proc with gopsutil", func() {
	// Start a dead loop so that current process has cpu usage.
	go func() {
		i := 1
		select {
		case <-time.After(10 * time.Second):
			return
		default:
			for i > 0 {
				i += 1
			}
		}
	}()
	testPid := int32(os.Getpid())
	procTestSource := &ProcPsutilSource{}
	Context("scenario cpu", func() {
		It("ProcPsutilSource is a correct ProcSource", func() {
			var iSource base.Source
			iSource = procTestSource
			_, ok := iSource.(base.ProcSource)
			Expect(ok).To(BeTrue())
		})
		It("process cpu should not be 0", func() {
			cpu, err := procTestSource.GetProcCpuPercent(testPid, 3*time.Second)
			Expect(err).NotTo(HaveOccurred())
			Expect(cpu.PitType).To(Equal(protos.EnumPitType_PROC_CPU_PERCENT))
			Expect(cpu.GetProcCpuPercent()).NotTo(BeNil())
			Expect(cpu.GetProcCpuPercent().Pid).To(Equal(int32(testPid)))
			Expect(cpu.GetProcCpuPercent().Percent == 0.0).To(BeFalse(), fmt.Sprintf("%d", testPid))
		})
	})
	Context("scenario mem info", func() {
		It("process mem should not be 0", func() {
			mem, err := procTestSource.GetProcMemInfo(testPid)
			Expect(err).NotTo(HaveOccurred())
			Expect(mem.GetProcMemInfo().Pid).To(Equal(int32(testPid)))
			Expect(mem.PitType).To(Equal(protos.EnumPitType_PROC_MEM_INFO))
			Expect(mem.GetProcMemInfo().RSS == 0).To(BeFalse())
		})
	})
	Context("scenario mem percent", func() {
		It("process mem should not be 0", func() {
			mem, err := procTestSource.GetProcMemPercent(testPid)
			Expect(err).NotTo(HaveOccurred())
			Expect(mem.GetProcMemPercent().Pid).To(Equal(int32(testPid)))
			Expect(mem.PitType).To(Equal(protos.EnumPitType_PROC_MEM_PERCENT))
			Expect(mem.GetProcMemPercent().Percent == 0.0).To(BeFalse())
		})
	})

})
