package financialmath

type OptionPriceModel interface {
	Init(underlying float64, strike float64, dte int, volatility float64, riskFreeInterest float64, dividend float64)
	Calc()
	Put() float64
	Call() float64
	SetVolatility(sigma float64)
}
