package strand

import "strings"

var tab = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRna converts DNA string to RNA (eliding non-DNA chars).
// E.g. ToRna("GxCTA") = "CGAU".
func ToRna(src string) string {
	f := func(r rune) rune {
		if repl, ok := tab[r]; ok {
			return repl
		}
		return -1
	}
	return strings.Map(f, src)
}
