package dft

import (
	"math"
	"math/cmplx"
)

func DFT(x []complex128) []complex128 {
	N := len(x)
	o := make([]complex128, N)

	if N <= 1 {
		copy(o, x)
		return o
	}

	for k := range o {
		for n, xn := range x {
			knN := float64(k) * float64(n) / float64(N)
			o[k] += xn * cmplx.Exp(complex(-2.0*math.Pi*knN, 0)*complex(0, 1))
		}
	}

	return o
}

func DFTReal(x []float64) []complex128 {
	return DFT(ToComplex(x))
}

// ToComplex returns the complex equivalent of the real-valued slice.
func ToComplex(x []float64) []complex128 {
	y := make([]complex128, len(x))
	for n, v := range x {
		y[n] = complex(v, 0)
	}
	return y
}
