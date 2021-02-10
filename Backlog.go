package main

// Backlog struct
// an accumulation of uncompleted work or matters needing to be dealt with.
// "the company took on extra staff to clear the backlog of work"
type Backlog struct {
}

// NewBacklog creates a Backlog struct.
func NewBacklog() *Backlog {
	return &Backlog{}
}

// Filter returns a filterd value.
func (p *Backlog) Filter(x float64) float64 {
	return x
}
