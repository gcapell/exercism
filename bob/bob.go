package bob

import "strings"

func Hey(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "Fine. Be that way!"
	}
	if s == strings.ToUpper(s) && s != strings.ToLower(s) {
		return "Whoa, chill out!"
	}
	if (s[len(s)-1]) == '?' {
		return "Sure."
	}
	return "Whatever."
}
