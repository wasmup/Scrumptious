package main

// Increment
// The Increment is the sum of all the Product Backlog items completed during a Sprint and the
// value of the increments of all previous Sprints. At the end of a Sprint, the new Increment must
// be “Done,” which means it must be in useable condition and meet the Scrum Team’s definition
// of “Done.” An increment is a body of inspectable, done work that supports empiricism at the
// end of the Sprint. The increment is a step toward a vision or goal. The increment must be in
// useable condition regardless of whether the Product Owner decides to release it.
type Increment struct {
}

// NewIncrement creates a Increment struct.
func NewIncrement() *Increment {
	return &Increment{}
}

// Filter returns a filterd value.
func (p *Increment) Filter(x float64) float64 {
	return x
}
