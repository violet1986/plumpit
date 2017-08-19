package systempit

import "plumpit/base"

type SystemPitGoGenerator func(base.Source, chan interface{})

func GetSystemGoGenerators(metricFuncs []base.SourceHandlerFunc) []SystemPitGoGenerator {
	result := []SystemPitGoGenerator{}
	for _, f := range metricFuncs {
		result = append(result, func(s base.Source, ch chan interface{}) {
			ch <- f(s)
		})
	}
	return result
}
