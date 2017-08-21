package base

import (
	"fmt"
	"plumpit/protos"
)

type SourceFunc func(Source) (protos.PitMessage, error)

type SourceDelegator func(...interface{}) SourceFunc

func delegatorTypeError(typeWanted string) error {
	return fmt.Errorf("Source is not a correct %s type", typeWanted)
}
func GetSystemCpuDelegator(...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if sysSource, ok := s.(SystemSource); ok {
			return sysSource.GetSystemCpu(), nil
		}
		return protos.PitMessage{}, delegatorTypeError("System Source")
	}
}

func GetProcCpuDelegator(args ...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if procSource, ok := s.(ProcSource); ok {
			if len(args) > 0 {
				return procSource.GetProcCpu(args[0]), nil
			}
		}
		return protos.PitMessage{}, delegatorTypeError("Process Source")
	}
}

func GetProcMemDelegator(args ...interface{}) SourceFunc {
	return func(s Source) (protos.PitMessage, error) {
		if procSource, ok := s.(ProcSource); ok {
			return procSource.GetProcMem(), nil
		}
		return protos.PitMessage{}, delegatorTypeError("Process Source")
	}
}
