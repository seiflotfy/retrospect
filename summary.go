package hindsight

import (
	"math"
	"time"
)

// Summary represents the stats on elapsed durations for a namespace
type Summary struct {
	namespace string
	last      time.Duration
	min       time.Duration
	max       time.Duration
	total     time.Duration
	avg       time.Duration
	count     uint64
}

func newSummary(namespace string) *Summary {
	return &Summary{
		namespace: namespace,
		min:       time.Duration(math.MaxInt64),
	}
}

func (s *Summary) update(elapsed time.Duration) {
	s.last = elapsed
	if s.min > elapsed {
		s.min = elapsed
	}
	if s.max < elapsed {
		s.max = elapsed
	}
	s.total += elapsed
	s.count++
	s.avg = s.total / time.Duration(s.count)
}

// Namespace returns the observation namespace
func (s *Summary) Namespace() string {
	return s.namespace
}

// Last returns the last elapsed duration
func (s *Summary) Last() time.Duration {
	return s.last
}

// Min returns the minimum elpased duration
func (s *Summary) Min() time.Duration {
	return s.min
}

// Max returns the maximum elapsed duration
func (s *Summary) Max() time.Duration {
	return s.max
}

// Total returns total elapsed duration
func (s *Summary) Total() time.Duration {
	return s.total
}

// Average returns the average elapsed duration
func (s *Summary) Average() time.Duration {
	return s.avg
}

// Count returns the number of elapsed measurements
func (s *Summary) Count() uint64 {
	return s.count
}
