package brent

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrent(t *testing.T) {
	a := assert.New(t)
	f := func(x float64) float64 {
		return x - 1
	}

	z, err := Brent(-10, 10, 0.001, f)
	if a.Nil(err) {
		a.Greater(0.001, math.Abs(1-z))
	}

	f2 := func(x float64) float64 {
		return (x + 3) * (x - 1) * (x - 1)
	}

	z, err = Brent(-4, -2, 0.001, f2)
	if a.Nil(err) {
		a.Greater(0.001, math.Abs(-3-z))
	}
}
