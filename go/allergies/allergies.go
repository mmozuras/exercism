package allergies

const testVersion = 1

var allergies = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(score uint) []string {
	result := []string{}
	for allergy, score := range allergies {
		if score&score > 0 {
			result = append(result, allergy)
		}
	}
	return result
}

func AllergicTo(score uint, allergen string) bool {
	return score&allergies[allergen] > 0
}
