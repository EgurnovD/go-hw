package common

import (
	"strconv"
	"sync"
)

type ReportMap map[string]int

type countersMap map[int]int

type Counters struct {
	mx sync.Mutex
	m  countersMap
}

const TotalKey string = "total"

// Summarize total and return a report
func (c *Counters) Report() ReportMap {
	c.mx.Lock()
	defer c.mx.Unlock()
	report := make(ReportMap)
	total := 0
	for k, v := range c.m {
		report[strconv.Itoa(k)] = v
		total += v
	}
	report[TotalKey] = total
	return report
}

// Increment key by 1
func (c *Counters) Inc(key int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

// Create new Counters
func NewCounters() *Counters {
	return &Counters{
		m: make(countersMap),
	}
}
