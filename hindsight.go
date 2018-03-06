package hindsight

import (
	"fmt"
	"math/rand"
	"time"
)

// Hindsight ...
type Hindsight struct {
	summary *Summary
	queue   chan *Observation
}

// New ...
func New(namespace string) *Hindsight {
	hs := &Hindsight{
		summary: newSummary(namespace),
		queue:   make(chan *Observation, 1e5),
	}
	go hs.dequeue()
	return hs
}

func (hs *Hindsight) dequeue() {
	for {
		select {
		case o := <-hs.queue:
			hs.summary.update(o.elapsed)
			o.count = hs.summary.count
			if o.done != nil {
				o.done(*o)
			}
		}
	}
}

func (hs *Hindsight) push(id uint64, elapsed time.Duration, done CallbackFunc, payload interface{}) {
	hs.queue <- &Observation{
		namespace: hs.summary.namespace,
		id:        id,
		elapsed:   elapsed,
		count:     0,
		done:      done,
		payload:   payload,
	}
}

// Observe ...
func (hs *Hindsight) Observe(done CallbackFunc, payload interface{}) (uint64, func(error) error) {
	return hs.ObserveSlow(0, done, payload)
}

// ObserveSlow ...
func (hs *Hindsight) ObserveSlow(maxDuration time.Duration, done CallbackFunc, payload interface{}) (uint64, func(error) error) {
	now := time.Now()
	id := rand.Uint64()
	finished := false
	return id, func(err error) error {
		if finished {
			return fmt.Errorf("observation over, can't call same func twice")
		}
		elapsed := time.Since(now)
		hs.push(id, elapsed, done, payload)
		finished = true
		return nil
	}
}

// Summary ...
func (hs *Hindsight) Summary() *Summary {
	return hs.summary
}
