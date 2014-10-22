package raindrops

import "strconv"

func Convert(n int) string {
	switch {
	case n%105 == 0:
		return "PlingPlangPlong"
	case n%35 == 0:
		return "PlangPlong"
	case n%15 == 0:
		return "PlingPlang"
	case n%21 == 0:
		return "PlingPlong"
	case n%3 == 0:
		return "Pling"
	case n%5 == 0:
		return "Plang"
	case n%7 == 0:
		return "Plong"
	}
	return strconv.Itoa(n)
}
