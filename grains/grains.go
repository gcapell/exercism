package grains

import "fmt"

func Square(n int) (uint64, error) {
	if n > 64 || n < 1 {
		return 0, fmt.Errorf("got %d, want 1-64", n)
	}
	return 1 << uint64(n-1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
