package main

import "time"

// Event struct for Scrum Events:
// Prescribed events are used in Scrum to create regularity and to minimize the need for meetings
// not defined in Scrum. All events are time-boxed events, such that every event has a maximum
// duration. Once a Sprint begins, its duration is fixed and cannot be shortened or lengthened. The
// remaining events may end whenever the purpose of the event is achieved, ensuring an
// appropriate amount of time is spent without allowing waste in the process.
// Other than the Sprint itself, which is a container for all other events, each event in Scrum is a
// formal opportunity to inspect and adapt something. These events are specifically designed to
// enable critical transparency and inspection. Failure to include any of these events results in
// reduced transparency and is a lost opportunity to inspect and adapt.
type Event struct {
	maxDuration time.Duration
}

// NewEvent creates new Scrum Event.
func NewEvent(maxDuration time.Duration) *Event {
	return &Event{maxDuration: maxDuration}
}

// Filter returns a filterd value.
func (p *Event) Filter(x float64) float64 {
	return x
}
