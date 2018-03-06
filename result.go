package hindsight

import "time"

// CallbackFunc ...
type CallbackFunc func(Result)

// Result ...
type Result struct {
	namespace string
	id        uint64
	elapsed   time.Duration
	count     uint64
	payload   interface{}

	done CallbackFunc
}

// Count ...
func (res *Result) Count() uint64 {
	return res.count
}

// Namespace ...
func (res *Result) Namespace() string {
	return res.namespace
}

// ID ...
func (res *Result) ID() uint64 {
	return res.id
}

// Elapsed ...
func (res *Result) Elapsed() time.Duration {
	return res.elapsed
}

// Payload ...
func (res *Result) Payload() interface{} {
	return res.payload
}
