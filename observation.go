package hindsight

import "time"

// CallbackFunc ...
type CallbackFunc func(Observation)

// Observation ...
type Observation struct {
	namespace string
	id        uint64
	elapsed   time.Duration
	count     uint64
	payload   interface{}

	done CallbackFunc
}

// Count ...
func (ob *Observation) Count() uint64 {
	return ob.count
}

// Namespace ...
func (ob *Observation) Namespace() string {
	return ob.namespace
}

// ID ...
func (ob *Observation) ID() uint64 {
	return ob.id
}

// Elapsed ...
func (ob *Observation) Elapsed() time.Duration {
	return ob.elapsed
}

// Payload ...
func (ob *Observation) Payload() interface{} {
	return ob.payload
}
