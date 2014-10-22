package cryptosquare

import (
	"bytes"
	"math"
	"strings"
	"unicode"
)

func Encode(plain string) string {
	normalised := strings.Map(normal, plain)
	side := int(math.Ceil(math.Sqrt(float64(len(normalised)))))
	var reply bytes.Buffer
	for j := 0; j < side*side; j++ {
		col, row := j/side, j%side
		tx := row*side + col
		if tx >= len(normalised) {
			continue
		}
		if reply.Len()%6 == 5 {
			reply.WriteByte(' ')
		}
		reply.WriteByte(normalised[tx])
	}
	return reply.String()
}

func normal(r rune) rune {
	if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
		return -1
	}
	return unicode.ToLower(r)
}
