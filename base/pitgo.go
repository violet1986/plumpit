package base

type PitGoGenerator func(Source, chan interface{})

func GoGenerators(funcs []SourceHandlerFunc) []PitGoGenerator {
	result := []PitGoGenerator{}
	for _, f := range funcs {
		result = append(result, func(s Source, ch chan interface{}) {
			ch <- f(s)
		})
	}
	return result
}
