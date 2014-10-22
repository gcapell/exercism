package scrabble_score

import (
	"strings"
)

var scores [27] int

func init() {
	data := map[int]string{
		1:  "aeioulnrst",
		2:  "dg",
		3:  "bcmp",
		4:  "fhvwy",
		5:  "k",
		8:  "jx",
		10: "qz",
	}
	for score, letters := range data {
		for _, r := range letters {
			scores[r-'a'] = score
		}
	}
}

func Score(s string) int {
	n := 0
	for _, r := range strings.ToLower(s) {
		d := r- 'a'
		if d < 0 || d>25 {
			continue
		}
		n += scores[d]
	}
	return n
}
