package financialmath

import (
	"math"

	binomialtree "github.com/gpayer/financial-math/binomial-tree"
	"github.com/gpayer/financial-math/brent"
)

func IV(isCall bool, optionPrice, strike, underlyingPrice float64, dte int, r float64) (float64, error) {
	return IVwithModel(&binomialtree.CoxRossRubinsteinTree{}, isCall, optionPrice, strike, underlyingPrice, dte, r)
}

func IVwithModel(m OptionPriceModel, isCall bool, optionPrice, strike, underlyingPrice float64, dte int, r float64) (float64, error) {
	m.Init(underlyingPrice, strike, dte, .01, r, 0)

	f := func(iv float64) float64 {
		m.SetVolatility(iv)
		m.Calc()
		var p float64
		if isCall {
			p = m.Call()
		} else {
			p = m.Put()
		}
		if math.IsNaN(p) {
			p = 0
		}
		// fmt.Printf("DEBUG: iv: %.2f, p - optionPrice: %.2f - %.2f = %.2f\n", iv, p, optionPrice, p-optionPrice)
		return p - optionPrice
	}

	return brent.Brent(.001, 10, 0.001, f)
}
