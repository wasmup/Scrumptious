package main

// Inspector struct
// Inspection
// Scrum users must frequently inspect Scrum artifacts and progress toward a Sprint Goal to detect
// undesirable variances. Their inspection should not be so frequent that inspection gets in the way
// of the work. Inspections are most beneficial when diligently performed by skilled inspectors at
// the point of work.
//
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
type Inspector struct {
}

// NewInspector creates a Inspector struct.
func NewInspector() *Inspector {
	return &Inspector{}
}

// Filter returns a filterd value.
func (p *Inspector) Filter(x float64) float64 {
	return x
}
