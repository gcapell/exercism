package triangle

import "math"

type Kind int

const (
	Equ Kind = iota
	Iso
	Sca
	NaT
)

var kindNames = []string{
	"Equ", "Iso", "Sca", "NaT",
}

func (k Kind) String() string {
	return kindNames[k]
}

func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a+b+c)   {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if !(a+b > c && a+c > b && b+c > a) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
