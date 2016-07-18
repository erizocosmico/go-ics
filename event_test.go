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
			ID:    "4",
			Start: d("20150830T193000Z"),
			End:   d("20150830T203000Z"),
		},
		{
			ID:           "3",
			Start:        d("20150930T123000Z"),
			End:          d("20150930T133000Z"),
			RecurrenceID: d("20150930T103000Z"),
		},
	}

	sort.Sort(byDate(eventList))
	result := ExcludeRecurrences(eventList)

	if len(result) != 4 {
		t.Errorf("Expected result length to be 3, not %d", len(result))
	}

	for i, r := range result {
		if (i < 2 || i == 3) && r.RecurrenceID.IsZero() {
			t.Errorf("Expected recurrent event with id %s, %v", r.ID, r)
		} else if i == 2 && !r.RecurrenceID.IsZero() {
			t.Errorf("Expected non-recurrent event with id %s, %v", r.ID, r)
		}
	}
}

func d(t string) time.Time {
	tm, _ := time.Parse(icsFormat, t)
	return tm
}
