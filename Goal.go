package main

// Goal struct
// The Sprint Goal is an objective set for the Sprint that can be met through the implementation of
// Product Backlog. It provides guidance to the Development Team on why it is building the
// Increment. It is created during the Sprint Planning meeting. The Sprint Goal gives the
// Development Team some flexibility regarding the functionality implemented within the Sprint.
// The selected Product Backlog items deliver one coherent function, which can be the Sprint Goal.
// The Sprint Goal can be any other coherence that causes the Development Team to work
// together rather than on separate initiatives.
//
// As the Development Team works, it keeps the Sprint Goal in mind. In order to satisfy the Sprint
// Goal, it implements functionality and technology. If the work turns out to be different than the
// Development Team expected, they collaborate with the Product Owner to negotiate the scope
// of Sprint Backlog within the Sprint.
type Goal struct {
}

// NewGoal creates a Goal struct.
func NewGoal() *Goal {
	return &Goal{}
}

// Filter returns a filterd value.
func (p *Goal) Filter(x float64) float64 {
	return x
}
