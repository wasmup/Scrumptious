package main

// Dev struct
// Only members of the Development Team create the Increment.
// Development Teams have the following characteristics:
// • They are self-organizing. No one (not even the Scrum Master) tells the Development Team
// how to turn Product Backlog into Increments of potentially releasable functionality;
// • Development Teams are cross-functional, with all the skills as a team necessary to create a
// product Increment;
// • Scrum recognizes no titles for Development Team members, regardless of the work being
// performed by the person;
// • Scrum recognizes no sub-teams in the Development Team, regardless of domains that need
// to be addressed like testing, architecture, operations, or business analysis; and,
// • Individual Development Team members may have specialized skills and areas of focus, but
// accountability belongs to the Development Team as a whole
type Dev struct {
}

// NewDev creates a Dev struct.
func NewDev() *Dev {
	return &Dev{}
}

// Filter returns a filterd value.
func (p *Dev) Filter(x float64) float64 {
	return x
}
