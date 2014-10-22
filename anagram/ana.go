package anagram

import "strings"

func GDetect(src string, candidates []string) []string {
	src = strings.ToLower(src)
	var counts [26]int
	for _, r := range src {
		counts[r-'a']++
	}
	reply := make([]string, 0, len(candidates))
	for _, c := range candidates {
		c = strings.ToLower(c)
		if sameCounts(c, counts) && c != src {
			reply = append(reply, c)
		}
	}
	return reply
}

func sameCounts(s string, c [26]int) bool {
	for _, r := range s {
		n := c[r-'a']
		if n < 1 {
			return false
		}
		c[r-'a'] = n - 1
	}
	for _, n := range c {
		if n > 0 {
			return false
		}
	}
	return true
}
