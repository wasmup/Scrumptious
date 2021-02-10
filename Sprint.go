package main

// Sprint struct
type Sprint struct {
}

// NewSprint creates a Sprint struct.
func NewSprint() *Sprint {
	return &Sprint{}
}

// Filter returns a filterd value.
func (p *Sprint) Filter(x float64) float64 {
	return x
}
