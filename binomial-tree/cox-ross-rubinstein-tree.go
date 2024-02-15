package binomialtree

import (
	"fmt"
	"math"
)

type CoxRossRubinsteinTree struct {
	Underlying       float64
	Strike           float64
	DTE              int
	Volatility       float64
	RiskFreeInterest float64
	Dividend         float64
	N                int
	timeToExpiration float64
	Delta            float64
}

func (c *CoxRossRubinsteinTree) SetVolatility(sigma float64) {
	c.Volatility = sigma
}

func (c *CoxRossRubinsteinTree) Init(underlying float64, strike float64, dte int, volatility float64, riskFreeInterest float64, dividend float64) {
	c.N = 1000
	c.Underlying = underlying
	c.Strike = strike
	c.DTE = dte
	c.Volatility = volatility
	c.RiskFreeInterest = riskFreeInterest
	c.Dividend = dividend

	c.Calc()
}

func (c *CoxRossRubinsteinTree) Calc() {
	c.timeToExpiration = float64(c.DTE) / 365
}

func (c *CoxRossRubinsteinTree) Put() float64 {
	deltaT := c.timeToExpiration / float64(c.N)
	up := math.Exp(c.Volatility * math.Sqrt(deltaT))
	p0 := (up*math.Exp(-c.Dividend*deltaT) - math.Exp(-c.RiskFreeInterest*deltaT)) / (up*up - 1)
	p1 := math.Exp(-c.RiskFreeInterest*deltaT) - p0

	p := make([]float64, c.N+1)

	for i := 0; i <= c.N; i++ {
		p[i] = c.Strike - c.Underlying*math.Pow(up, float64(2*i-c.N))
		if p[i] < 0 {
			p[i] = 0
		}
	}

	for j := c.N - 1; j >= 0; j-- {
		for i := 0; i <= j; i++ {
			p[i] = p0*p[i+1] + p1*p[i]
			exercise := math.Max(0, c.Strike-c.Underlying*math.Pow(up, float64(2*i-j)))
			if p[i] < exercise {
				p[i] = exercise
			}
		}

		// grab aproximation for delta
		if j == 1 {
			c.Delta = (p[1] - p[0]) / (c.Underlying*up - c.Underlying/up)
		}
	}

	if math.IsNaN(p[0]) {
		fmt.Println("ERROR: CoxRossRubinstein: NaN")
		return 0
	}
	return p[0]
}

func (c *CoxRossRubinsteinTree) Call() float64 {
	deltaT := c.timeToExpiration / float64(c.N)
	up := math.Exp(c.Volatility * math.Sqrt(deltaT))
	p0 := (up*math.Exp(-c.Dividend*deltaT) - math.Exp(-c.RiskFreeInterest*deltaT)) / (up*up - 1)
	p1 := math.Exp(-c.RiskFreeInterest*deltaT) - p0

	C := make([]float64, c.N+1)

	for i := 0; i <= c.N; i++ {
		C[i] = c.Underlying*math.Pow(up, float64(2*i-c.N)) - c.Strike
		if C[i] < 0 {
			C[i] = 0
		}
	}

	for j := c.N - 1; j >= 0; j-- {
		for i := 0; i <= j; i++ {
			C[i] = p0*C[i+1] + p1*C[i]
			exercise := math.Max(0, c.Underlying*math.Pow(up, float64(2*i-j))-c.Strike)
			if C[i] < exercise {
				C[i] = exercise
			}
		}

		// grab aproximation for delta
		if j == 1 {
			c.Delta = (C[1] - C[0]) / (c.Underlying*up - c.Underlying/up)
		}
	}

	if math.IsNaN(C[0]) {
		fmt.Println("ERROR: CoxRossRubinstein: NaN")
		return 0
	}
	return C[0]
}
