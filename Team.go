package main

// Team struct
type Team struct {
}

// NewTeam creates a Team struct.
func NewTeam() *Team {
	return &Team{}
}

// Filter returns a filterd value.
func (p *Team) Filter(x float64) float64 {
	return x
}
