package queenattack

import (
	"errors"
	"fmt"
)

var (
	errSameSquare = errors.New("same square")
)

func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errSameSquare
	}
	wp, err := parse(w)
	if err != nil {
		return false, err
	}
	bp, err := parse(b)
	if err != nil {
		return false, err
	}
	return wp.x == bp.x || wp.y == bp.y || abs(wp.x-bp.x) == abs(wp.y-bp.y), nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type square struct{ x, y int }

func parse(s string) (square, error) {
	var reply square
	if len(s) != 2 {
		return reply, fmt.Errorf("parse(%s) wrong length", s)
	}
	reply.x = int(s[0]) - int('a')
	reply.y = int(s[1]) - int('1')

	if reply.x < 0 || reply.x > 7 || reply.y < 0 || reply.y > 7 {
		return reply, fmt.Errorf("parse(%s) out-of-bounds", s)
	}
	return reply, nil
}
