package test

import (
	m "../packages"
	"testing"
)

type testpair struct {
	values []float64
	max    float64
	min    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 2, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1},
	{[]float64{-1, 1}, 1, -1},
}

func TestMath(t *testing.T) {
	for _, pair := range tests {
		max := m.Max(pair.values)
		if max != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", max)
		}

		min := m.Min(pair.values)
		if min != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", min)
		}
	}
}
