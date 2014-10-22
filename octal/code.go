package octal

import "fmt"

func ParseOctal(s string) (int64, error) {
	var reply int64
	for _, r := range s {
		if r < '0' || r > '7' {
			return 0, fmt.Errorf("bad char %q in %q", r, s)
		}
		reply = reply*8 + int64(r-'0')
	}
	return reply, nil
}
