package say

const testVersion = 1

var translations = map[uint64]string{
	0:    "zero",
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	1e2:  "hundred",
	1e3:  "thousand",
	1e6:  "million",
	1e9:  "billion",
	1e12: "trillion",
	1e15: "quadrillion",
	1e18: "quintillion",
}

var checks = []uint64{1e2, 1e3, 1e6, 1e9, 1e12, 1e15, 1e18}

func Say(input uint64) string {
	if input < 20 {
		return translations[input]
	} else if input < 100 {
		remainder := input % 10
		result := translations[input-remainder]
		if remainder > 0 {
			result = result + "-" + Say(remainder)
		}
		return result
	}
	for i, check := range checks {
		if input < check {
			return toEnglish(input, checks[i-1])
		}
	}
	return toEnglish(input, 1e18)
}

func toEnglish(number uint64, decimal uint64) string {
	remainder := number % decimal
	result := Say(number/decimal) + " " + translations[decimal]
	if remainder > 0 {
		result = result + " " + Say(remainder)
	}
	return result
}
