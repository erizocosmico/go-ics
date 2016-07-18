package ics

import (
	"sort"
	"time"
)

// Event represents an event in the calendar
type Event struct {
	Start         time.Time
	End           time.Time
	Created       time.Time
	Modified      time.Time
	AlarmTime     time.Duration
	ID            string
	Status        string
	Description   string
	Location      string
	Summary       string
	RRule         string
	RecurrenceID  time.Time
	Class         string
	Sequence      int
	Attendees     []Attendee
	Organizer     Attendee
	WholeDayEvent bool
}

type byDate []Event

func (e byDate) Len() int {
	return len(e)
}

func (e byDate) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e byDate) Less(i, j int) bool {
	return e[i].Start.Before(e[j].Start)
}

// NewEvent returns a new empty Event entity
func NewEvent() *Event {
	return &Event{
		Attendees: []Attendee{},
	}
}

// Clone returns an identical clone of the current Event entity
func (e *Event) Clone() *Event {
	newEvent := *e
	return &newEvent
}

func (e *Event) Equals(e2 *Event) bool {
	return e.Start.Equal(e2.Start) && e.End.Equal(e2.End) && e.Summary == e2.Summary
}

// ExcludeRecurrences receives a list of events and removes the repetitions that
// have been overriden
func ExcludeRecurrences(evs []Event) []Event {
	result := []Event{}
	eventsByID := make(map[string][]Event)
	for _, e := range evs {
		if _, ok := eventsByID[e.ID]; !ok {
			eventsByID[e.ID] = []Event{e}
		} else {
			eventsByID[e.ID] = append(eventsByID[e.ID], e)
		}
	}

	for _, evs := range eventsByID {
		if len(evs) == 1 {
			result = append(result, evs[0])
			continue
		}

		for i := 0; i < len(evs); i++ {
			if i+1 >= len(evs) {
				result = append(result, evs[i])
				continue
			}

			event := evs[i]
			nextEvent := evs[i+1]

			if event.ID == nextEvent.ID {
				if event.RecurrenceID.Equal(nextEvent.Start) {
					i++
				} else if nextEvent.RecurrenceID.Equal(event.Start) {
					i++
					result = append(result, nextEvent)
					continue
				}
			}

			result = append(result, event)
		}
	}

	sort.Sort(byDate(result))
	return result
}
