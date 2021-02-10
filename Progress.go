package main

// Progress struct
// Monitoring Progress Toward Goals
// At any point in time, the total work remaining to reach a goal can be summed. The Product
// Owner tracks this total work remaining at least every Sprint Review. The Product Owner
// compares this amount with work remaining at previous Sprint Reviews to assess progress
// toward completing projected work by the desired time for the goal. This information is made
// transparent to all stakeholders.
// Various projective practices upon trending have been used to forecast progress, like burndowns, burn-ups, or cumulative flows. These have proven useful. However, these do not replace
// the importance of empiricism. In complex environments, what will happen is unknown. Only
// what has already happened may be used for forward-looking decision-making.
//
// Monitoring Sprint Progress
// At any point in time in a Sprint, the total work remaining in the Sprint Backlog can be summed.
// The Development Team tracks this total work remaining at least for every Daily Scrum to project
// the likelihood of achieving the Sprint Goal. By tracking the remaining work throughout the
// Sprint, the Development Team can manage its progress
type Progress struct {
}

// NewProgress creates a Progress struct.
func NewProgress() *Progress {
	return &Progress{}
}

// Filter returns a filterd value.
func (p *Progress) Filter(x float64) float64 {
	return x
}
