package main

import (
    "fmt"
    "math"
)

type Stats struct {
    sigma       float64
    sigma_sq    float64
    count       int
}

func (this *Stats) StatsInit() {
    this.sigma = 0.0
    this.sigma_sq = 0.0
    this.count = 0
}

func (this *Stats) StatsCount() int {
    return this.count
}

func (this *Stats) StatsAccumulate(fval float64) {
    this.sigma += fval
    this.sigma_sq += fval * fval
    this.count++
}

func (this *Stats) StatsMean() float64 {
    if 0 == this.count { return this.sigma }
    return this.sigma / float64(this.count)
}

func (this *Stats) StatsStdev(ddof int) float64 {
    if this.count < 2 { return 0 }
    M := (float64(this.count) * this.sigma_sq) - (this.sigma * this.sigma)
    variance := M / float64(this.count * (this.count - ddof))
    return math.Sqrt(variance)
}

func (this *Stats) StatsStdevS() float64 {
    return this.StatsStdev(1)
}

func (this *Stats) StatsStdevP() float64 {
    return this.StatsStdev(0)
}

func (this *Stats) StatsToString() string {
  return fmt.Sprintf("count=%d mean=%0.3g stdev_s=%0.3g",
    this.StatsCount(),
    this.StatsMean(),
    this.StatsStdevS() )
}

func main() {
  var s Stats
  s.StatsInit()

  data := [...]float64 {2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7} 
  
  for _, num := range data {
    s.StatsAccumulate(num)
  }

  fmt.Println(s.StatsToString())
}
