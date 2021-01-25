// Package stats implements some basic descriptive statistics functions.
package stats

import "math"

// Context is the struct holding the accumulated statistics.
type Context struct {
	sigma   float64
	sigmaSq float64
	count   int
}

const ci95 float64 = 1.960

// New creates and returns a new stats context.
func New() *Context {
	return new(Context)
}

// Reset clears/resets the stats context.
func (ctx *Context) Reset() {
	ctx.sigma = 0.0
	ctx.sigmaSq = 0.0
	ctx.count = 0
}

// Count returns a count of the number of values accumulated.
func (ctx *Context) Count() int {
	return ctx.count
}

func (ctx *Context) Add(vals ...float64) {
	for _, v := range vals {
		ctx.sigma += v
		// ctx.sigmaSq += v * v
		ctx.sigmaSq = math.FMA(v, v, ctx.sigmaSq)
		ctx.count++
	}
}

func (ctx *Context) Sum() float64 {
	return ctx.sigma
}

func (ctx *Context) Mean() float64 {
	if 0 == ctx.count {
		return 0.0
	}
	return ctx.sigma / float64(ctx.count)
}

func (ctx *Context) Stdev(ddof int) float64 {
	if ctx.count < 2 {
		return 0.0
	}

	m := (float64(ctx.count) * ctx.sigmaSq) - (ctx.sigma * ctx.sigma)
	variance := m / float64(ctx.count*(ctx.count-ddof))
	return math.Sqrt(variance)
}

func (ctx *Context) StdevS() float64 {
	return ctx.Stdev(1)
}

func (ctx *Context) StdevP() float64 {
	return ctx.Stdev(0)
}

func (ctx *Context) Ci95() float64 {
	if ctx.count < 2 {
		return 0.0
	}
	return ci95 * (ctx.StdevS() / math.Sqrt(float64(ctx.count)))
}
