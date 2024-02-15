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

func TestDelta(t *testing.T) {
	a := assert.New(t)
	coxtree := &CoxRossRubinsteinTree{}
	coxtree.Init(100, 100, 1, 0.01, 0.00, 0.00)
	coxtree.Calc()
	_ = coxtree.Put()
	a.Greater(0.005, math.Abs(-0.5-coxtree.Delta))
	_ = coxtree.Call()
	a.Greater(0.005, math.Abs(0.5-coxtree.Delta))
}

func TestDelta2(t *testing.T) {
	a := assert.New(t)
	coxtree := &CoxRossRubinsteinTree{}
	coxtree.Init(100, 1, 1, 0.01, 0.00, 0.00)
	coxtree.Calc()
	_ = coxtree.Put()
	a.Equal(0.0, coxtree.Delta)
	_ = coxtree.Call()
	a.Greater(0.005, math.Abs(1.0-coxtree.Delta))
}

func TestDelta3(t *testing.T) {
	a := assert.New(t)
	coxtree := &CoxRossRubinsteinTree{}
	coxtree.Init(100, 200, 1, 0.01, 0.00, 0.00)
	coxtree.Calc()
	_ = coxtree.Put()
	a.Equal(-1.0, coxtree.Delta)
	_ = coxtree.Call()
	a.Equal(0.0, coxtree.Delta)
}
