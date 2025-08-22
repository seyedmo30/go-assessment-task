package service

import (
	"assessment/domain/entity"
	"sort"
	"time"
)

// TODO : append unit test

// event represents a point in time where the current usage changes.
// + at start -> +quantity
// + at end   -> -quantity
type event struct {
	t      time.Time
	change int64
}

// buildEvents converts a slice of plannings into a slice of events.
// Each planning generates two events: +quantity at StartAt and -quantity at EndAt.
// The returned slice is sorted by time (ascending). For equal times,
// events with negative change come before positive change so we free resources before allocating.
func buildEvents(plannings []entity.Planning) []event {
	events := make([]event, 0, len(plannings)*2)

	for _, p := range plannings {
		events = append(events, event{t: p.StartAt, change: p.Quantity})
		events = append(events, event{t: p.EndAt, change: -p.Quantity})
	}

	// Comparator explained:
	// 1) earlier time should come first
	// 2) if times are equal, negative changes (releases) should come before positive changes (allocations)
	sort.Slice(events, func(i, j int) bool {
		if events[i].t.Before(events[j].t) {
			return true
		}
		if events[j].t.Before(events[i].t) {
			return false
		}
		// same timestamp => release (-) before allocate (+)
		return events[i].change < events[j].change
	})

	return events
}

// maxConcurrentUsage runs a simple sweep over sorted events and returns the maximum concurrent usage.
// It accumulates changes and tracks the highest value seen.
func maxConcurrentUsage(events []event) int64 {
	var concurrent int64
	var maxConcurrent int64
	for _, e := range events {
		concurrent += e.change
		if concurrent > maxConcurrent {
			maxConcurrent = concurrent
		}
	}
	return maxConcurrent
}
