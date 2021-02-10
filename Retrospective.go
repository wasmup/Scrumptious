package main

// Retrospective struct
type Retrospective struct {
}

// NewRetrospective creates a Retrospective struct.
func NewRetrospective() *Retrospective {
	return &Retrospective{}
}

// Filter returns a filterd value.
func (p *Retrospective) Filter(x float64) float64 {
	return x
}
