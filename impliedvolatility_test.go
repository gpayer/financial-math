package financialmath

import (
	"math"
	"testing"

	binomialtree "github.com/gpayer/financial-math/binomial-tree"
	"github.com/stretchr/testify/assert"
)

func TestIV(t *testing.T) {
	a := assert.New(t)
	ivCall, err := IV(true, 1.72, 40, 39.15, 35, .03)
	if a.Nil(err) {
		// fmt.Printf("iv value: %f\n", ivCall)
		a.GreaterOrEqual(0.001, math.Abs(ivCall-0.423))
	}

	ivPut, err := IV(false, 2.58, 40, 39.15, 35, .03)
	if a.Nil(err) {
		a.GreaterOrEqual(0.001, math.Abs(ivPut-0.447))
	}

	_, err = IV(true, 2.214, 24, 26.21, 16, .01)

	ivCall, err = IVwithModel(&binomialtree.CoxRossRubinsteinTree{}, true, 2.34, 20, 22.00, 32, .03)
	if a.Nil(err) {
		a.GreaterOrEqual(0.001, math.Abs(ivCall-0.401))
		// fmt.Printf("iv value: %f\n", ivCall)
	}

	a.Nil(err)
}
