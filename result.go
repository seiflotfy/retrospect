package retrospect

import "time"

// CallbackFunc ...
type CallbackFunc func(Result)

// Result contains the data about the measurement/observation
type Result struct {
	namespace string
	elapsed   time.Duration
	count     uint64
	payload   interface{}

	done CallbackFunc
}

// Count returns the number of the  measurment
func (res *Result) Count() uint64 {
	return res.count
}

// Namespace returns the observer namespace
func (res *Result) Namespace() string {
	return res.namespace
}

// Elapsed returns the elapsed duration of the measurement
func (res *Result) Elapsed() time.Duration {
	return res.elapsed
}

// Payload the user passed data to the measurment/observation
func (res *Result) Payload() interface{} {
	return res.payload
}
