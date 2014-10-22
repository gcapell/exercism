package robotname

import "fmt"

// Robot tracks serial numbers
type Robot string

// Name (optionally initialises and) returns robot's name.
func (r *Robot) Name() string {
	if *r == "" {
		r.Reset()
	}
	return string(*r)
}

var maxLetters, maxID int

func name() string {
	maxID++
	if maxID == 1000 {
		maxID = 0
		maxLetters++
	}
	return fmt.Sprintf("%c%c%03d", 'A'+maxLetters/26, 'A'+maxLetters%26, maxID)
}

// Reset chooses a new (unique) name.
func (r *Robot) Reset() {
	*r = Robot(name())
}
