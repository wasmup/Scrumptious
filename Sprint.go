package main

// Sprint struct
// The Sprint
// The heart of Scrum is a Sprint, a time-box of one month or less during which a “Done”, useable,
// and potentially releasable product Increment is created. Sprints have consistent durations
// throughout a development effort. A new Sprint starts immediately after the conclusion of the
// previous Sprint.
// Sprints contain and consist of the Sprint Planning, Daily Scrums, the development work, the
// Sprint Review, and the Sprint Retrospective.
// During the Sprint:
// • No changes are made that would endanger the Sprint Goal;
// • Quality goals do not decrease; and,
// • Scope may be clarified and re-negotiated between the Product Owner and Development
// Team as more is learned.
// Each Sprint may be considered a project with no more than a one-month horizon. Like projects,
// Sprints are used to accomplish something. Each Sprint has a goal of what is to be built, a design
// and flexible plan that will guide building it, the work, and the resultant product increment.
// Sprints are limited to one calendar month. When a Sprint’s horizon is too long the definition of
// what is being built may change, complexity may rise, and risk may increase. Sprints enable
// predictability by ensuring inspection and adaptation of progress toward a Sprint Goal at least
// every calendar month. Sprints also limit risk to one calendar month of cost.
type Sprint struct {
	goal string
	work string
	Planning
	ds []DailyScrum
	Review
	Retrospective
	inspectors []Inspector
}

// NewSprint creates a Sprint struct.
func NewSprint() *Sprint {
	return &Sprint{}
}

// Filter returns a filterd value.
func (p *Sprint) Filter(x float64) float64 {
	return x
}
