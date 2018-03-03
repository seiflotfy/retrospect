package observer

import (
	"time"
)

var instance *Report

type callBackFunc func(id string, sum Summary, payload ...interface{})

// Init ...
func Init(namespace string) {
	instance = &Report{
		namespace: namespace,
		summaries: make(map[string]*Summary),
	}
}

// This ...
func This(id string, callback callBackFunc) {
}

// ThisOnSlowt ...
func ThisOnSlowt(id string, tolerance time.Duration, callback callBackFunc) {
	This(id, callback)
}

// PrintThis ...
func PrintThis(id string) {
}

// PrintThisOnSlowt ...
func PrintThisOnSlow(id string, tolerance time.Duration) {
}
