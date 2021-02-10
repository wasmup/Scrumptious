package main

// Adaptation struct
// Adaptation
// If an inspector determines that one or more aspects of a process deviate outside acceptable
// limits, and that the resulting product will be unacceptable, the process or the material being
// processed must be adjusted. An adjustment must be made as soon as possible to minimize
// further deviation.
//
// Scrum prescribes four formal events for inspection and adaptation, as described in the Scrum
// Events section of this document:
// • Sprint Planning
// • Daily Scrum
// • Sprint Review
// • Sprint Retrospective
type Adaptation struct {
}

// NewAdaptation creates a Adaptation struct.
func NewAdaptation() *Adaptation {
	return &Adaptation{}
}

// Filter returns a filterd value.
func (p *Adaptation) Filter(x float64) float64 {
	return x
}
