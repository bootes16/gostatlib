package stats

import "math"

// Accumulated statistics
// sum of values
// sum of square of values
// number of values
type context struct {
	sigma   float64
	sigmaSq float64
	count   int
}

func New() *context {
	return new(context)
}

func (ctx *context) Reset() {
	ctx.sigma = 0.0
	ctx.sigmaSq = 0.0
	ctx.count = 0
}

func (ctx *context) GetCount() int {
	return ctx.count
}

func (ctx *context) Add(vals ...float64) {
	for _, v := range vals {
		ctx.sigma += v
		ctx.sigmaSq += v * v
		ctx.count++
	}
}

func (ctx *context) Sum() float64 {
	return ctx.sigma
}

func (ctx *context) Mean() float64 {
	if 0 == ctx.count {
		return 0.0
	}
	return ctx.sigma / float64(ctx.count)
}

func (ctx *context) Stdev(ddof int) float64 {
	if ctx.count < 2 {
		return 0.0
	}

	m := (float64(ctx.count) * ctx.sigmaSq) - (ctx.sigma * ctx.sigma)
	variance := m / float64(ctx.count*(ctx.count-ddof))
	return math.Sqrt(variance)
}

func (ctx *context) StdevS() float64 {
	return ctx.Stdev(1)
}

func (ctx *context) StdevP() float64 {
	return ctx.Stdev(0)
}
