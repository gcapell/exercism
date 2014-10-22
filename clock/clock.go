package clock

import (
	"fmt"
)

type Clock struct {
	h, m int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(m int) Clock {
	m += c.m
	h, m := c.h+m/60, m%60
	if m < 0 {
		h, m = h-1, m+60
	}
	h %= 24
	if h < 0 {
		h += 24
	}
	return Clock{h, m}
}

func New(h, m int) Clock {
	return Clock{h, m}
}
