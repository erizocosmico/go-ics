package ics

import (
	"sort"
	"testing"
	"time"
)

func TestExcludeRecurrences(t *testing.T) {
	eventList := []Event{
		{
			ID:    "1",
			Start: d("20150830T103000Z"),
			End:   d("20150830T123000Z"),
		},
		{
			ID:    "2",
			Start: d("20150830T093000Z"),
			End:   d("20150830T103000Z"),
		},
		{
			ID:           "1",
			Start:        d("20150830T113000Z"),
			End:          d("20150830T123000Z"),
			RecurrenceID: d("20150830T103000Z"),
		},
		{
			ID:           "2",
			Start:        d("20150830T123000Z"),
			End:          d("20150830T133000Z"),
			RecurrenceID: d("20150830T093000Z"),
		},
		{
			ID:           "3",
			Start:        d("20150930T123000Z"),
			End:          d("20150930T133000Z"),
			RecurrenceID: d("20150930T103000Z"),
		},
	}

	sort.Sort(events(eventList))
	result := ExcludeRecurrences(eventList)

	if len(result) != 3 {
		t.Errorf("Expected result length to be 3, not %d", len(result))
	}

	for _, r := range result {
		if r.RecurrenceID.IsZero() {
			t.Errorf("Expected recurrent event with id %s, %v", r.ID, r)
		}
	}
}

func d(t string) time.Time {
	tm, _ := time.Parse(icsFormat, t)
	return tm
}
