package hindsight

import (
	"math"
	"time"
)

// Summary ...
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

// Namespace ...
func (s *Summary) Namespace() string {
	return s.namespace
}

// Last ..
func (s *Summary) Last() time.Duration {
	return s.last
}

// Min ...
func (s *Summary) Min() time.Duration {
	return s.min
}

// Max ...
func (s *Summary) Max() time.Duration {
	return s.max
}

// Total ...
func (s *Summary) Total() time.Duration {
	return s.total
}

// Average ...
func (s *Summary) Average() time.Duration {
	return s.avg
}

// Count ...
func (s *Summary) Count() uint64 {
	return s.count
}
