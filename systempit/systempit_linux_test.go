var _ = Describe("Tests only run in Linux", func() {
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
		It("process cpu should not be 0", func() {
			cpu, err := procTestSource.GetProcCpuPercent(testPid, 3*time.Second)
			Expect(err).NotTo(HaveOccurred())
			Expect(cpu.PitType).To(Equal(protos.EnumPitType_PROC_CPU_PERCENT))
			Expect(cpu.GetProcCpuPercent()).NotTo(BeNil())
			Expect(cpu.GetProcCpuPercent().Pid).To(Equal(int32(testPid)))
			Expect(cpu.GetProcCpuPercent().Percent == 0.0).To(BeFalse(), fmt.Sprintf("%d", testPid))
		})
	})
})