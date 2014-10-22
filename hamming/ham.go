package hamming

func Distance(a, b string) int {
	var d int
	min := len(a)
	if len(b) < min {
		min = len(b)
	}
	for j := 0; j < min; j++ {
		if a[j] != b[j] {
			d++
		}
	}
	return d
}
