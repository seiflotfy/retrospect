package observer

import "time"

type Report struct {
	namespace string
	summaries map[string]*Summary
}

type Summary struct {
	namespace string
	id        string
	last      time.Duration
	min       time.Duration
	max       time.Duration
	avg       time.Duration
	count     uint64
	tags      map[string]interface{}
}
