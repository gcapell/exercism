package pythagorean

import (
	"math"
	"sort"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	var reply []Triplet
	maxM := int(math.Sqrt(float64(max)))
	for m := 2; m < maxM; m++ {
		for n := 1 + m%2; n < m; n += 2 {
			if !coprime(n, m) {
				continue
			}
			m2, n2 := m*m, n*n
			a, b, c := m2-n2, 2*m*n, m2+n2
			for k := 1; ; k++ {
				if c*k > max {
					break
				}
				if a*k < min {
					continue
				}
				reply = append(reply, Triplet{k * a, k * b, k * c})
			}
		}
	}
	sort.Sort(ByLex(reply))
	return reply
}

func Sum(p int) []Triplet {
	var reply []Triplet
	maxM := int(math.Sqrt(float64(p)))
	for m := 2; m < maxM; m++ {
		for n := 1 + m%2; n < m; n += 2 {
			if !coprime(n, m) {
				continue
			}
			m2, n2 := m*m, n*n
			a, b, c := m2-n2, 2*m*n, m2+n2
			sum := a + b + c
			if p%sum == 0 {
				k := p / sum
				if a > b {
					a, b = b, a
				}
				reply = append(reply, Triplet{k * a, k * b, k * c})
			}
		}
	}
	sort.Sort(ByLex(reply))
	return reply
}

type ByLex []Triplet

func (ts ByLex) Len() int           { return len(ts) }
func (ts ByLex) Swap(i, j int)      { ts[i], ts[j] = ts[j], ts[i] }
func (ts ByLex) Less(i, j int) bool { return ts[i][0] < ts[j][0] }

func coprime(a, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}
	return a == 1
}
