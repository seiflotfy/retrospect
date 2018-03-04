package observer

import "time"

// ReportFunc ...
type ReportFunc func(id string, duration time.Duration, count uint64)

// Report ...
type Report struct {
	id    string
	last  time.Duration
	min   time.Duration
	max   time.Duration
	total time.Duration
	avg   time.Duration
	count uint64
}

func (r *Report) update(elapsed time.Duration) {
	r.last = elapsed
	if r.min > elapsed {
		r.min = elapsed
	}
	if r.max < elapsed {
		r.max = elapsed
	}
	r.total += elapsed
	r.count++
	r.avg = r.total / time.Duration(r.count)
}
