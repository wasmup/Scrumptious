package main

import "time"

// Planning struct
// Sprint Planning
// The work to be performed in the Sprint is planned at the Sprint Planning. This plan is created by
// the collaborative work of the entire Scrum Team.
//
// Sprint Planning is time-boxed to a maximum of eight hours for a one-month Sprint. For shorter
// Sprints, the event is usually shorter. The Scrum Master ensures that the event takes place and
// that attendants understand its purpose. The Scrum Master teaches the Scrum Team to keep it
// within the time-box.
//
// The input to this meeting is the Product Backlog, the latest product Increment, projected
// capacity of the Development Team during the Sprint, and past performance of the Development
// Team. The number of items selected from the Product Backlog for the Sprint is solely up to the
// Development Team. Only the Development Team can assess what it can accomplish over the
// upcoming Sprint.
//
// The Development Team usually starts by designing the system and the work needed to convert
// the Product Backlog into a working product Increment. Work may be of varying size, or
// estimated effort. However, enough work is planned during Sprint Planning for the Development
// Team to forecast what it believes it can do in the upcoming Sprint. Work planned for the first
// days of the Sprint by the Development Team is decomposed by the end of this meeting, often to
// units of one day or less. The Development Team self-organizes to undertake the work in the
// Sprint Backlog, both during Sprint Planning and as needed throughout the Sprint.
//
// The Product Owner can help to clarify the selected Product Backlog items and make trade-offs.
// If the Development Team determines it has too much or too little work, it may renegotiate the
// selected Product Backlog items with the Product Owner. The Development Team may also invite
// other people to attend to provide technical or domain advice
//
// By the end of the Sprint Planning, the Development Team should be able to explain to the
// Product Owner and Scrum Master how it intends to work as a self-organizing team to
// accomplish the Sprint Goal and create the anticipated Increment.
type Planning struct {
	// max eight hours for a one-month Sprint.
	Duration time.Duration
	// Sprint Planning answers the following:
	// • What can be delivered in the Increment resulting from the upcoming Sprint?
	// • How will the work needed to deliver the Increment be achieved?
	Sprint
}

// NewPlanning creates a Planning struct.
func NewPlanning() *Planning {
	return &Planning{}
}

// Filter returns a filterd value.
func (p *Planning) Filter(x float64) float64 {
	return x
}
