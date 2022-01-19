package trig

import (
	"math"
)

// Versine functions
func Versin(θ float64) float64 {
	return 1. - math.Cos(θ)
}

func Coversin(θ float64) float64 {
	return 1. - math.Sin(θ)
}

func Vercosin(θ float64) float64 {
	return 1. + math.Cos(θ)
}

func Covercosin(θ float64) float64 {
	return 1. + math.Sin(θ)
}

func Haversin(θ float64) float64 {
	return (1. - math.Cos(θ)) / 2.
}

func Havercosin(θ float64) float64 {
	return (1. + math.Cos(θ)) / 2.
}

func Hacoversin(θ float64) float64 {
	return (1. - math.Sin(θ)) / 2.
}

func Hacovercosin(θ float64) float64 {
	return (1. + math.Sin(θ)) / 2.
}

// Inverse Versine functions
func Aversin(y float64) float64 {
	return math.Acos(1. - y)
}

func Avercos(y float64) float64 {
	return math.Acos(1. + y)
}

func Acoversin(y float64) float64 {
	return math.Asin(1. - y)
}

func Acovercos(y float64) float64 {
	return math.Asin(1. + y)
}

func Ahaversin(y float64) float64 {
	return 2. * math.Asin(math.Sqrt(y))
}

func Ahavercos(y float64) float64 {
	return 2. * math.Acos(math.Sqrt(y))
}

// functions not in standard lib
func Sec(θ float64) float64 {
	return 1. / math.Cos(θ)
}

func Csc(θ float64) float64 {
	return 1. / math.Sin(θ)
}

func Cot(θ float64) float64 {
	return 1. / math.Tan(θ)
}

func Sech(θ float64) float64 {
	return 1. / math.Cosh(θ)
}

func Exsec(θ float64) float64 {
	return Sec(θ) - 1.
}

func Excsc(θ float64) float64 {
	return Csc(θ) - 1.
}

func Chord(θ float64) float64 {
	return 2. * math.Sin(θ/2.)
}
