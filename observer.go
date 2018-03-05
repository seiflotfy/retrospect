package observer

import (
	"fmt"
	"math"
	"time"
)

// Observer ...
type Observer struct {
	namespace string
	reports   map[string]*Report
}

// New ...
func New(namespace string) *Observer {
	return &Observer{
		namespace: namespace,
		reports:   make(map[string]*Report),
	}
}

func (obs *Observer) update(id string, elapsed time.Duration) {
	if _, ok := obs.reports[id]; !ok {
		obs.reports[id] = &Report{
			id:  id,
			min: time.Duration(math.MaxInt64),
		}
	}
	obs.reports[id].update(elapsed)
}

// Clear ...
func (obs *Observer) Clear(id string) {
	delete(obs.reports, id)
}

// ClearAll ...
func (obs *Observer) ClearAll(id string) {
	obs.reports = make(map[string]*Report)
}

// Measure ...
func (obs *Observer) Measure(id string, callback ReportFunc) func() {
	return obs.MeasureOnSlow(id, 0, callback)
}

// MeasureOnSlow ...
func (obs *Observer) MeasureOnSlow(id string, maxDuration time.Duration, callback ReportFunc) func() {
	now := time.Now()
	return func() {
		elapsed := time.Since(now)
		obs.update(id, elapsed)
		if callback != nil {
			callback(id, elapsed, obs.reports[id].count)
		}
	}
}

// Get ...
func (obs *Observer) Get(id string) *Report {
	fmt.Println(obs.reports)
	return obs.reports[id]
}
