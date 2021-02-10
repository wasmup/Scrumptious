package main

// Impediment struct
// Scrum Master Service to the Development Team: Removing impediments
type Impediment struct {
}

// NewImpediment creates a Impediment struct.
func NewImpediment() *Impediment {
	return &Impediment{}
}

// Filter returns a filterd value.
func (p *Impediment) Filter(x float64) float64 {
	return x
}
