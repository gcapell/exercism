package luhn

import "strconv"

func Valid(s string) bool {
	ok, sum := luhnSum(s, false)
	return ok && sum == 0
}

func AddCheck(s string) string {
	_, sum := luhnSum(s, true)
	add := 0
	if sum != 0 {
		add = 10 - sum
	}
	return s + digits[add]
}

var digits = map[int]string{}

func init() {
	for j := 0; j < 10; j++ {
		digits[j] = strconv.Itoa(j)
	}
}

func luhnSum(s string, double bool) (bool, int) {
	sum := 0
	for j := len(s) - 1; j >= 0; j-- {
		c := s[j]
		if !(c >= '0' && c <= '9') {
			continue
		}
		n := int(c - '0')
		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		double = !double
		sum += n
	}
	return sum > 0, sum % 10
}
