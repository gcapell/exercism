package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	reply := make(map[string]int)
	for score, letters := range in {
		for _, letter := range letters {
			reply[strings.ToLower(letter)] = score
		}
	}
	return reply
}
