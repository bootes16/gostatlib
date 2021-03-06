package main

import (
	"math"
	"testing"

	"../mathx/stats"
)

func compareFloat(a float64, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestStatsCreate(t *testing.T) {
	s := stats.New()
	s.Reset()
}

func TestStatsSum(t *testing.T) {
	s := stats.New()
	s.Add(1.0, 2.0, 3.0)
	sum := s.Sum()
	if sum != 6.0 {
		t.Errorf("expected: %f received: %f", 6.0, sum)
	}
}
func TestStatsMean(t *testing.T) {
	s := stats.New()
	s.Add(1.0, 2.0, 3.0)
	mean := s.Mean()
	if mean != 2.0 {
		t.Errorf("expected: %f received: %f", 2.0, mean)
	}
}

func TestStatsStdDev(t *testing.T) {
	s := stats.New()
	s.Add(1.0, 2.0, 3.0)
	stdev := s.StdevS()
	if 1 != stdev {
		t.Errorf("expected: %f received: %f", 1.0, stdev)
	}

	s.Reset()
	s.Add(2.0, 4.0, 6.0)
	stdev = s.StdevS()
	if 2 != stdev {
		t.Errorf("expected: %f received: %f", 2.0, stdev)
	}

	popStdev := s.StdevP()
	delta := math.Abs(popStdev - 1.633)
	if delta > 0.0005 {
		t.Errorf("expected: %f received: %f", 1.633, popStdev)
	}
}

func TestStatsCI(t *testing.T) {
	s := stats.New()
	s.Add(25.0, 35.0, 10.0, 17.0, 29.0, 14.0, 21.0, 31.0)
	ci := s.Ci95()
	if !compareFloat(ci, 6.0723, 0.0001) {
		t.Errorf("Confidence interval is %f", ci)
	}
}
