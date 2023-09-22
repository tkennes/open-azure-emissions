package test_helpers

import (
	"math"
)

const float64EqualityThreshold = 1e-9

func FloatsAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
