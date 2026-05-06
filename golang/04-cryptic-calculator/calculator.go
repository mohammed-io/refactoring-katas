package kata

import "math"

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) normalize(a []float64, b float64, c1 float64) []float64 {
	d := []float64{}
	e := math.Inf(1)
	f := math.Inf(-1)
	for g := 0; g < len(a); g++ {
		if a[g] < e {
			e = a[g]
		}
		if a[g] > f {
			f = a[g]
		}
	}
	h := f - e
	i := c1 - b
	for g := 0; g < len(a); g++ {
		k := b + ((a[g]-e)/h)*i
		d = append(d, math.Round(k*100)/100)
	}
	return d
}
