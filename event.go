package ics

import "time"

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
	Class         string
	Sequence      int
	Attendees     []Attendee
	Organizer     Attendee
	WholeDayEvent bool
}

type events []Event

func (e events) Len() int {
	return len(e)
}

func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e events) Less(i, j int) bool {
	return e[j].Start.Before(e[i].Start)
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
