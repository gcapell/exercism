
        package palindrome

import "errors"

type Product struct {
	Product int // palindromic, of course
	// list of all possible two-factor factorizations of Product, within
	// given limits, in order
	Factorizations [][2]int
}

func (p *Product) update(product, j, k int, max bool) {
	if product == p.Product {
		p.Factorizations = append(p.Factorizations, [2]int{j, k})
		return
	}
	if p.Product != 0 {
		switch max {
		case true:
			if product < p.Product {
				return
			}
		case false:
			if product > p.Product {
				return
			}
		}
	}
	p.Product = product
	p.Factorizations = [][2]int{[2]int{j, k}}
}

var errNoPalindromes = errors.New("No palindromes...")
var errOrder = errors.New("fmin > fmax...")

func Products(fmin, fmax int) (pmin, pmax Product, e error) {
	if fmin>fmax {
		e = errOrder
		return
	}
	for j := fmin; j <= fmax; j++ {
		for k := j; k <= fmax; k++ {
			product := j * k
			if !isPal(product) {
				continue
			}
			pmin.update(product, j, k, false)
			pmax.update(product, j, k, true)
		}
	}
	if pmin.Product == 0 {
		e = errNoPalindromes
	}
	return
}

func isPal(n int) bool {
	var digits []byte
	for ; n > 0; n /= 10 {
		digits = append(digits, byte(n%10))
	}
	for j, k := 0, len(digits)-1; j < k; j, k = j+1, k-1 {
		if digits[j] != digits[k] {
			return false
		}
	}
	return true
}
      
