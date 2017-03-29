package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	} else if isEquilateral(a, b, c) {
		return Equ
	} else if isIsoceles(a, b, c) {
		return Iso
	} else if isScalene(a, b, c) {
		return Sca
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	return isSide(a) && isSide(b) && isSide(c) &&
		a+b >= c && a+c >= b && b+c >= a
}

func isSide(f float64) bool {
	return f > 0 && !math.IsNaN(f) && !math.IsInf(f, 1)
}

func isEquilateral(a, b, c float64) bool {
	return a == b && a == c && b == c
}

func isIsoceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

func isScalene(a, b, c float64) bool {
	return a != b && a != c && b != c
}

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)
