package ics

import (
	"testing"
	"time"
)

func TestCalendarInfo(t *testing.T) {
	calendar, err := ParseCalendar("testCalendars/2eventsCal.ics", 0, nil)
	if err != nil {
		t.Errorf("Failed to parse the calendar ( %s ) \n", err.Error())
	}

	if calendar.Name != "2 Events Cal" {
		t.Errorf("Expected name '%s' calendar , got '%s' calendars \n", "2 Events Cal", calendar.Name)
	}

	if calendar.Description != "The cal has 2 events(1st with attendees and second without)" {
		t.Errorf("Expected description '%s' calendar , got '%s' calendars \n", "The cal has 2 events(1st with attendees and second without)", calendar.Description)
	}

	if calendar.Version != 2.0 {
		t.Errorf("Expected version %f calendar, got %f\n", 2.0, calendar.Version)
	}

	events := calendar.Events
	if len(events) != 2 {
		t.Errorf("Expected %d events in calendar, got %d events\n", 2, len(events))
	}
}

func TestCalendarEvents(t *testing.T) {
	calendar, err := ParseCalendar("testCalendars/2eventsCal.ics", 1000, nil)
	if err != nil {
		t.Errorf("Failed to parse the calendar ( %s ) \n", err.Error())
	}
	tz, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		t.FailNow()
	}

	event := calendar.Events[0]
	start := time.Date(2014, time.Month(6), 16, 6, 0, 0, 0, tz)
	end := time.Date(2014, time.Month(6), 16, 7, 0, 0, 0, tz)
	created, _ := time.Parse(icsFormat, "20140515T075711Z")
	modified, _ := time.Parse(icsFormat, "20141125T074253Z")
	location := "In The Office"
	desc := "1. Report on previous weekly tasks. \\n2. Plan of the present weekly tasks."
	seq := 1
	status := "CONFIRMED"
	summary := "General Operative Meeting"
	rrule := ""
	attendeesCount := 3

	if !event.Start.Equal(start) {
		t.Errorf("Expected start %s, found %s\n", start, event.Start)
	}

	if !event.End.Equal(end) {
		t.Errorf("Expected end %s, found %s\n", end, event.End)
	}

	if event.Created != created {
		t.Errorf("Expected created %s, found %s\n", created, event.Created)
	}

	if event.Modified != modified {
		t.Errorf("Expected modified %s, found %s\n", modified, event.Modified)
	}

	if event.Location != location {
		t.Errorf("Expected location %s, found %s\n", location, event.Location)
	}

	if event.Description != desc {
		t.Errorf("Expected description %s, found %s\n", desc, event.Description)
	}

	if event.Sequence != seq {
		t.Errorf("Expected sequence %d, found %d\n", seq, event.Sequence)
	}

	if event.Status != status {
		t.Errorf("Expected status %s, found %s\n", status, event.Status)
	}

	if event.Summary != summary {
		t.Errorf("Expected status %s, found %s\n", summary, event.Summary)
	}

	if event.RRule != rrule {
		t.Errorf("Expected rrule %s, found %s\n", rrule, event.RRule)
	}

	if len(event.Attendees) != attendeesCount {
		t.Errorf("Expected attendeesCount %d, found %d\n", attendeesCount, len(event.Attendees))
	}
}

func TestParseEventDate(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		t.FailNow()
	}

	expected := time.Date(2015, time.Month(9), 30, 15, 0, 0, 0, loc)
	dataStart := "DTSTART;TZID=Europe/Madrid:20150930T150000\n"
	result, err := parseEventDate("DTSTART", dataStart)
	if err != nil {
		t.FailNow()
	}

	if !expected.Equal(result) {
		t.Errorf("Expected time %v to be %v", result, expected)
	}

	dataEnd := "DTEND;TZID=Europe/Madrid:20150930T150000\n"
	result, err = parseEventDate("DTEND", dataEnd)
	if err != nil {
		t.FailNow()
	}

	if !expected.Equal(result) {
		t.Errorf("Expected time %v to be %v", result, expected)
	}
}

func TestParseEventRecurrenceID(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		t.FailNow()
	}
	expected := time.Date(2015, time.Month(10), 13, 15, 0, 0, 0, loc)
	data := "RECURRENCE-ID;TZID=Europe/Madrid:20151013T150000\n"

	result, err := parseEventRecurrenceID(data)
	if err != nil {
		t.Error(err)
	}

	if !expected.Equal(result) {
		t.Errorf("Expected time %v to be %v", result, expected)
	}
}
