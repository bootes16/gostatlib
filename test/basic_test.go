package main

import (
	"main/mathx/stats"
	"math"
	"testing"
)

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
