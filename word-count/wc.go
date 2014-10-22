package wc

import (
	"strings"
	"unicode"
)

func WordCount(s string) Histogram {
	m := make(Histogram)

	for _, f := range strings.FieldsFunc(s, ignore) {
		m[strings.ToLower(f)]++
	}
	return m
}

func ignore(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}
