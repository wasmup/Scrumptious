package main

// Scope struct
type Scope struct {
}

// NewScope creates a Scope struct.
func NewScope() *Scope {
	return &Scope{}
}

// Filter returns a filterd value.
func (p *Scope) Filter(x float64) float64 {
	return x
}
