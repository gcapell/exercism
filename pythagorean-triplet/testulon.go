
        package pythagorean

type Triplet [3]int

func isTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if isTriplet(a, b, c) {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

func Sum(p int) []Triplet {
	triplets := []Triplet{}
	for a := 1; a <= p; a++ {
		for b := a; b <= p; b++ {
			if b <= p-a-b && isTriplet(a, b, p-a-b) {
				triplets = append(triplets, Triplet{a, b, p - a - b})
			}
		}
	}
	return triplets
}
      
