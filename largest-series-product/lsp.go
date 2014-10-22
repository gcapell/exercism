package lsproduct

import (
	"fmt"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	if len(digits) < span {
		return 0, fmt.Errorf("requested span %d from %d digits", span, len(digits))
	}

	max := 0
	chunks := strings.Split(digits, "0")
	for _, chunk := range chunks {
		p := lsp(chunk, span)
		if p > max {
			max = p
		}
	}
	return max, nil
}

type RingBuf struct {
	digits []int
	pos int
	product int
}

func (r *RingBuf) update(n int) int {
	r.product /= r.digits[r.pos]
	r.product *= n
	r.digits[r.pos] = n
	r.pos = (r.pos + 1) % len(r.digits)
	return r.product
}

func NewBuf(digits string) *RingBuf {
	ints := make([]int, 0, len(digits))
	product := 1
	for _, r := range digits {
		n :=  int(r - '0')
		ints = append(ints, n)
		product *= n
	}
	return &RingBuf{ints, 0, product}
}

func lsp(digits string, span int) int {
	if len(digits) < span {
		return 0
	}
	buf := NewBuf(digits[:span])
	
	max := buf.product
	for _, r := range digits[span:] {
		p := buf.update(int(r-'0'))
		if p > max {
			max = p
		}
	}
	return max
}
