package main

// Planning struct
// Sprint Planning is time-boxed to a maximum of eight hours for a one-month Sprint. For shorter
// Sprints, the event is usually shorter. The Scrum Master ensures that the event takes place and
// that attendants understand its purpose. The Scrum Master teaches the Scrum Team to keep it
// within the time-box.
type Planning struct {
	// Sprint Planning answers the following:
	// • What can be delivered in the Increment resulting from the upcoming Sprint?
	product string
	// • How will the work needed to deliver the Increment be achieved?
	workNeeded string
}

// NewPlanning creates a Planning struct.
func NewPlanning() *Planning {
	return &Planning{}
}

// Filter returns a filterd value.
func (p *Planning) Filter(x float64) float64 {
	return x
}
