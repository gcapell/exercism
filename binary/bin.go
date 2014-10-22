package binary

import "fmt"

// ParseBinary converts a string like '1011' into integer 11.
func ParseBinary(s string) (int, error) {
	var reply int
	for _, c := range s {
		reply <<= 1
		switch c {
		case '1':
			reply++
		case '0':
		default:
			return 0, fmt.Errorf("bad char %q in %q", c, s)
		}
	}
	return reply, nil
}
