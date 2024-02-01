package binomialtree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrices(t *testing.T) {
	a := assert.New(t)
	coxtree := &CoxRossRubinsteinTree{}
	coxtree.Init(39.15, 40, 35, 0.441, 0.01, 0.03)
	c := coxtree.Call()
	//fmt.Printf("call: %f\n", c)
	a.Greater(0.005, math.Abs(1.72-c))
	p := coxtree.Put()
	//fmt.Printf("put: %f\n", p)
	a.Greater(0.005, math.Abs(2.64-p))
}
