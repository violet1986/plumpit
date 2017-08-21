package base

type SourceHandlerFunc func(...interface{}) func(Source) interface{}

func GetSystemCPUHandler(...interface{}) func(Source) interface{} {
	return func(s Source) interface{} {
		if sysSource, ok := s.(SystemSource); ok {
			return sysSource.GetSystemCPU()
		}
		return nil
	}
}

func GetProcCPUHandler(args ...interface{}) func(Source) interface{} {
	return func(s Source) interface{} {
		if procSource, ok := s.(ProcSource); ok {
			return procSource.GetProcCPU(args[0])
		}
		return nil

	}
}
