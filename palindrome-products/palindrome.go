package palindrome

import (
	"errors"
	"math"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, e error) {
	if fmin > fmax {
		e = errors.New("fmin > fmax...")
		return
	}
	gen := palLess(fmax * fmax)
	for {
		pal := gen()
		if pal < fmin*fmin {
			e = errors.New("No palindromes...")
			return
		}
		if pmax.factors(pal, fmin, fmax) {
			break
		}
	}
	gen = palGreater(fmin * fmin)
	for {
		pal := gen()
		if pal == pmax.Product {
			pmin = pmax
			break
		}
		if pmin.factors(pal, fmin, fmax) {
			break
		}
	}
	return
}

// palLess returns function to generate palindromes <= n in descending order.
func palLess(n int) func() int {
	digits := int(math.Floor(math.Log10(float64(n)))) + 1
	rightDigits := digits / 2
	shift := int(math.Pow10(rightDigits))
	left := n / shift
	decade := int(math.Pow10(digits - rightDigits - 1))

	gen := func() int {
		reply := left*shift + reverse(left, digits)
		left--
		if left < decade {
			digits--
			rightDigits = digits / 2
			shift = int(math.Pow10(rightDigits))
			max := int(math.Pow10(digits - rightDigits))
			left = max - 1
			decade = max / 10
		}
		return reply
	}

	// skip past first one if too big
	if left*shift+reverse(left, digits) > n {
		gen()
	}

	return gen
}

// palLess returns function to generate palindromes >=  n in ascending order.
func palGreater(n int) func() int {
	digits := int(math.Floor(math.Log10(float64(n)))) + 1
	rightDigits := digits / 2
	shift := int(math.Pow10(rightDigits))
	left := n / shift
	decade := int(math.Pow10(digits - rightDigits))

	gen := func() int {

		reply := left*shift + reverse(left, digits)
		left++
		if left == decade {
			digits++
			rightDigits = digits / 2
			shift = int(math.Pow10(rightDigits))
			max := int(math.Pow10(digits - rightDigits - 1))
			left = max
			decade = max * 10
		}
		return reply
	}

	// skip past first one if too small
	if left*shift+reverse(left, digits) < n {
		gen()
	}

	return gen
}

// factors stores factorizations of product between fmin and fmax, or returns false.
func (p *Product) factors(product, fmin, fmax int) bool {
	fs, ok := primeFactors(product)
	if !ok {
		return false
	}

	groups := group(fs, 1, 1, fmin, fmax)
	if len(groups) == 0 {
		return false
	}
	p.Product = product
	p.Factorizations = groups

	return true
}

type ByFirst [][2]int

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i][0] < a[j][0] }

func group(factors []Factor, left, right, fmin, fmax int) [][2]int {
	if left > fmax || right > fmax {
		return nil
	}
	if len(factors) == 0 {
		if left > right {
			return nil
		}
		if left < fmin || right < fmin {
			return nil
		}
		return [][2]int{{left, right}}
	}

	f, tail := factors[0], factors[1:]
	var reply [][2]int

	products := make([]int, 0, f.count+1)
	p := 1
	for j := 0; j <= f.count; j++ {
		products = append(products, p)
		p *= f.prime
	}
	for j := 0; j < len(products); j++ {
		g := group(tail, left*products[j], right*products[len(products)-j-1], fmin, fmax)
		if len(g) > 0 {
			reply = append(reply, g...)
		}
	}

	return reply
}

func reverse(n, digits int) int {
	if digits%2 == 1 {
		n /= 10
	}
	var reply int
	for ; n > 0; n /= 10 {
		reply = reply*10 + n%10
	}
	return reply
}

var primes []int

func init() {
	primes = primesLessThan(1000)
}

type Factor struct{ prime, count int }

func primeFactors(n int) ([]Factor, bool) {
	var reply []Factor
	for _, p := range primes {
		var count int
		for n%p == 0 {
			count++
			n /= p
		}
		if count > 0 {
			reply = append(reply, Factor{p, count})
		}
		if n == 1 {
			return reply, true
		}
	}
	return nil, false
}

func primesLessThan(n int) []int {
	switch n {
	case 0, 1, 2:
		return []int{2}
	case 3:
		return []int{2, 3}
	}

	primes := make([]int, 0, 170)
	primes = append(primes, 2, 3)
	sieve := make([]bool, n)
	p := 5
	for p <= n {
		if !sieve[p] {
			primes = append(primes, p)
			filter(sieve, p)
		}
		p += 2
		if !sieve[p] {
			primes = append(primes, p)
			filter(sieve, p)
		}
		p += 4
	}
	// We might have added an extra
	if primes[len(primes)-1] > n {
		primes = primes[:len(primes)-1]
	}
	return primes
}

func filter(sieve []bool, n int) {
	for j := n; j < len(sieve); j += n {
		sieve[j] = true
	}
}
