package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.

func ConcurrentFrequency(l []string) FreqMap {
	m := FreqMap{}
	var wg sync.WaitGroup
	wg.Add(len(l))

	ch := make(chan FreqMap, len(l))
	for _, str := range l {
		go func(s string) {
			ch <- Frequency(s)
			wg.Done()
		}(str)
	}

	wg.Wait()
	close(ch)

	for fm := range ch {
		for k, v := range fm {
			m[k] += v
		}
	}
	return m
}
