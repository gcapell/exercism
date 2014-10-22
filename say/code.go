package say

import "strings"

var (
	digits    = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teens     = [...]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens      = [...]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	thousands = [...]string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "hexillion"}
)

func underThousand(n uint64) string {
	if n < 10 {
		return digits[n-1]
	}
	if n < 20 {
		return teens[n-10]
	}
	if n < 100 {
		return join(tens[n/10-2], "-", n%10)
	}
	return join(digits[n/100-1]+" hundred", " ", n%100)
}

func join(a, sep string, b uint64) string {
	if b > 0 {
		return a + sep + underThousand(b)
	}
	return a
}

// Say returns verbal representation of n.
func Say(n uint64) string {
	if n == 0 {
		return "zero"
	}
	var reply []string
	for _, s := range thousands {
		var rem uint64
		n, rem = n/1000, n%1000
		if rem > 0 {
			ns := underThousand(rem)
			if len(s) > 0 {
				ns += " " + s
			}
			reply = append(reply, ns)
		}
		if n == 0 {
			break
		}
	}
	for j, k := 0, len(reply)-1; j < k; j, k = j+1, k-1 {
		reply[j], reply[k] = reply[k], reply[j]
	}
	return strings.Join(reply, " ")
}
