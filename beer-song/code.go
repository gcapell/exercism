package beer

import (
	"fmt"
	"strings"
)

var special = map[int]string{
	0: "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n",
	1: "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n",
	2: "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n",
}

func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", fmt.Errorf("%d expected to be in range 0,99", n)
	}
	if s, ok := special[n]; ok {
		return s, nil
	}
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
}

func Verses(upper, lower int) (string, error) {
	if upper <= lower {
		return "", fmt.Errorf("verses upper %d <= %d", upper, lower)
	}
	var reply []string
	for j := upper; j >= lower; j-- {
		v, err := Verse(j)
		if err != nil {
			return "", err
		}
		reply = append(reply, v)
	}
	return strings.Join(reply, "\n") + "\n", nil
}

func Song() string {
	reply, err := Verses(99, 0)
	if err != nil {
		panic(err)
	}
	return reply
}
