package romannumerals

import "fmt"

var table = []struct {
	decade int
	tx     []string
}{
	{1000, []string{"M", "MM", "MMM"}},
	{100, []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}},
	{10, []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}},
	{1, []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}},
}

func ToRomanNumeral(n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("ToRomanNumeral(%d) too small", n)
	}
	if n > 3000 {
		return "", fmt.Errorf("ToRomanNumeral(%d) too big", n)
	}
	reply := ""
	for _, entry := range table {
		var this int
		this, n = n/entry.decade, n%entry.decade
		if this != 0 {
			reply += entry.tx[this-1]
		}
	}
	return reply, nil
}
