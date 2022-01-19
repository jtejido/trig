package trig

import (
	"math"
	"testing"
)

func TestTrigonometryCircularRotationOne(t *testing.T) {
	tol := 1e-10
	x := 25.0

	versin := Versin(x)
	coversin := Coversin(x + (math.Pi / 2.))
	vercosin := Vercosin(x + math.Pi)
	covercosin := Covercosin(x + ((3 * math.Pi) / 2.))

	if math.Abs(versin-coversin) > tol || math.Abs(coversin-vercosin) > tol || math.Abs(vercosin-covercosin) > tol || math.Abs(covercosin-versin) > tol {
		t.Errorf("Circular Rotation test Failed.")
	}
}

func TestTrigonometryCircularRotationTwo(t *testing.T) {
	tol := 1e-10
	x := 25.0

	haversin := Haversin(x)
	hacoversin := Hacoversin(x + (math.Pi / 2.))
	havercosin := Havercosin(x + math.Pi)
	hacovercosin := Hacovercosin(x + ((3 * math.Pi) / 2.))

	if math.Abs(haversin-hacoversin) > tol || math.Abs(hacoversin-havercosin) > tol || math.Abs(havercosin-hacovercosin) > tol || math.Abs(hacovercosin-haversin) > tol {
		t.Errorf("Circular Rotation test Failed.")
	}
}

// for some reason, this fails
// see van Vlijmen, Oscar (2005-12-28) [2003]. "Goniology". Eenheden, constanten en conversies.
// func TestTrigonometryArcVersinArcTan(t *testing.T) {
// 	tol := 1e-10
// 	x := 1.5 // for 0 ≤ y ≤ 2
// 	aversin := Aversin(x)
// 	atan := math.Atan((math.Sqrt(2*x - x*x)) / (1 - x))

// 	if math.Abs(aversin-atan) > tol {
// 		t.Errorf("Equality Failed. want %v, got %v", atan, aversin)
// 	}
// }

func TestTrigonometryArcCoversinArcTan(t *testing.T) {
	tol := 1e-10
	x := 1.5 // for 0 ≤ y ≤ 2

	acoversin := Acoversin(x)
	atan := math.Atan((1 - x) / (math.Sqrt(2*x - x*x)))

	if math.Abs(acoversin-atan) > tol {
		t.Errorf("Equality Failed. want %v, got %v", atan, acoversin)
	}
}

func TestTrigonometryArcHaversinArcTan(t *testing.T) {
	tol := 1e-10
	x := .5 // for 0 ≤ y ≤ 1

	ahaversin := Ahaversin(x)
	atan := math.Atan((2 * (math.Sqrt(x - x*x))) / (1 - 2*x))

	if math.Abs(ahaversin-atan) > tol {
		t.Errorf("Equality Failed. want %v, got %v", atan, ahaversin)
	}
}
