package letter

func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}

	ch := make(chan FreqMap)
	for _, s := range ss {
		go func(s string) { ch <- Frequency(s) }(s)
	}
	for _ = range ss {
		for r, v := range <-ch {
			m[r] += v
		}
	}
	return m
}
