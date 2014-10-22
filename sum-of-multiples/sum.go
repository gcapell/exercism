package summultiples

import "sort"

type Factor struct{ div, sign int }

func MultipleSummer(divs ...int) func(int) int {
	divs = removeMultiples(divs)

	var factors []Factor
	sign := 1
	for p := range divs {
		for _, div := range lcmOfN(divs, p+1) {
			factors = append(factors, Factor{div, sign})
		}
		sign *= -1
	}

	return func(x int) int {
		sum := 0
		for _, f := range factors {
			n := (x - 1) / f.div
			sum += f.sign * n * (n + 1) / 2 * f.div
		}
		return sum
	}
}

func removeMultiples(n []int) []int {
	sort.Ints(n)
	for j := 0; j < len(n); j++ {
		for k := j + 1; k < len(n); k++ {
			if n[k]%n[j] == 0 {
				copy(n[k:], n[k+1:])
				n = n[:len(n)-1]
			}
		}
	}
	return n
}

func lcm(a, b int) int {
	product := a * b
	// euclid gcd
	for b != 0 {
		a, b = b, a%b
	}
	return product / a
}

// lcmOfN returns a slice with the lcm of each combination of n elements of nums.
// e.g. lcmOfN( [3,4,5], 2) -> [ lcm(3,4), lcm(3,5), lcm(4,5)]
func lcmOfN(nums []int, n int) []int {
	if n == 1 {
		return nums
	}
	var reply []int
	for pos, x := range nums {
		for _, y := range lcmOfN(nums[pos+1:], n-1) {
			reply = append(reply, lcm(y, x))
		}
	}
	return reply
}
