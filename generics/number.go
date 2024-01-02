package generics

import (
	"golang.org/x/exp/constraints"
)

func countInts(arr []int) int {
	var s int
	for _, v := range arr {
		s += v
	}
	return s
}

func countFloats(arr []float64) float64 {
	var s float64
	for _, v := range arr {
		s += v
	}
	return s
}

type number interface {
	constraints.Integer | constraints.Float
}

func countNumbers[T number](arr []T) T {
	var s T
	for _, v := range arr {
		s += v
	}
	return s
}
