package ics

import "time"

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
