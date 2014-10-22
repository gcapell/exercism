package diffsquares

func SquareOfSums(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}

func SumOfSquares(n int) int {
	s := 0
	for j := 1; j <= n; j++ {
		s += j * j
	}
	return s
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
