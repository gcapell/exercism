package allergies

var maskToAllergen = []struct {
	mask     int
	allergen string
}{
	{1, "eggs"},
	{2, "peanuts"},
	{4, "shellfish"},
	{8, "strawberries"},
	{16, "tomatoes"},
	{32, "chocolate"},
	{64, "pollen"},
	{128, "cats"},
}

var allergenToMask = make(map[string]int, len(maskToAllergen))

func init() {
	for _, a := range maskToAllergen {
		allergenToMask[a.allergen] = a.mask
	}
}

func Allergies(score int) []string {
	reply := make([]string, 0, len(maskToAllergen))
	for _, a := range maskToAllergen {
		if score&a.mask != 0 {
			reply = append(reply, a.allergen)
		}
	}
	return reply
}

func AllergicTo(score int, allergen string) bool {
	mask, ok := allergenToMask[allergen]
	if !ok {
		return false // should be error
	}
	return score&mask != 0
}
