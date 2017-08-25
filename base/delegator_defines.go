package base

import (
	"fmt"
	"plumpit/protos"
)

type Unmarshaller func([]byte) (RawMessage, error)

type SourceFunc func(Source) (protos.PitMessage, error)

type SourceDelegator func(...interface{}) SourceFunc

func delegatorTypeError(typeWanted string) error {
	return fmt.Errorf("Source is not a correct %s type", typeWanted)
}
func GetSystemCpuDelegator(...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if sysSource, ok := s.(SystemSource); ok {
			return sysSource.GetSystemCpu()
		}
		return protos.PitMessage{}, delegatorTypeError("System Source")
	}
}

func GetProcCpuPercentDelegator(args ...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if procSource, ok := s.(ProcSource); ok {
			if len(args) > 0 {
				return procSource.GetProcCpuPercent(args[0])
			}
		}
		return protos.PitMessage{}, delegatorTypeError("Process Source")
	}
}

func GetProcMemInfoDelegator(args ...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if procSource, ok := s.(ProcSource); ok {
			return procSource.GetProcMemInfo()
		}
		return protos.PitMessage{}, delegatorTypeError("Process Source")
	}
}

func GetProcMemPercentDelegator(args ...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if procSource, ok := s.(ProcSource); ok {
			return procSource.GetProcMemPercent()
		}
		return protos.PitMessage{}, delegatorTypeError("Process Source")
	}
}
