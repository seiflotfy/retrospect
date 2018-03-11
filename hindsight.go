package retrospect

import (
	"fmt"
	"time"
)

// Retrospect ...
type Retrospect struct {
	summary *Summary
	queue   chan *Result
	clear   chan struct{}
	stop    chan struct{}
}

// New ...
func New(namespace string) *Retrospect {
	hs := &Retrospect{
		summary: newSummary(namespace),
		queue:   make(chan *Result, 1e5),
	}
	go hs.dequeue()
	return hs
}

func (hs *Retrospect) dequeue() {
	for {
		select {
		case o := <-hs.queue:
			hs.summary.update(o.elapsed)
			o.count = hs.summary.count
			if o.done != nil {
				go o.done(*o)
			}
		case <-hs.clear:
			hs.queue = make(chan *Result)
		case <-hs.stop:
			return
		}
	}
}

func (hs *Retrospect) push(elapsed time.Duration, done CallbackFunc, payload interface{}) {
	hs.queue <- &Result{
		namespace: hs.summary.namespace,
		elapsed:   elapsed,
		count:     0,
		done:      done,
		payload:   payload,
	}
}

// Observe returns a function that when called measures the elapsed duration
// and triggeres the done with the payload as an argument
// if called again it will return error
func (hs *Retrospect) Observe(done CallbackFunc, payload interface{}) func() error {
	now := time.Now()
	finished := false
	return func() error {
		if finished {
			return fmt.Errorf("observation over, can't call same func twice")
		}
		elapsed := time.Since(now)
		hs.push(elapsed, done, payload)
		finished = true
		return nil
	}
}

// Summary return a summary of all Results
func (hs *Retrospect) Summary() *Summary {
	return hs.summary
}

// Clear purges all observations
func (hs *Retrospect) Clear() {
	hs.clear <- struct{}{}
}

// Stop all observations and basically kill the observer
func (hs *Retrospect) Stop() {
	hs.stop <- struct{}{}
}
