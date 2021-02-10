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
	// During Sprint Planning the Scrum Team also crafts a Sprint Goal. The Sprint Goal is an objective
	// that will be met within the Sprint through the implementation of the Product Backlog, and it
	// provides guidance to the Development Team on why it is building the Increment.
	Goals []Goal
	// 	Definition of “Done”
	// When a Product Backlog item or an Increment is described as “Done”, everyone must
	// understand what “Done” means. Although this may vary significantly per Scrum Team, members
	// must have a shared understanding of what it means for work to be complete, to ensure
	// transparency. This is the definition of “Done” for the Scrum Team and is used to assess when
	// work is complete on the product Increment.
	// The same definition guides the Development Team in knowing how many Product Backlog items
	// it can select during a Sprint Planning. The purpose of each Sprint is to deliver Increments of
	// potentially releasable functionality that adhere to the Scrum Team’s current definition of
	// “Done.”
	// Development Teams deliver an Increment of product functionality every Sprint. This Increment
	// is useable, so a Product Owner may choose to immediately release it. If the definition of "Done"
	// for an increment is part of the conventions, standards or guidelines of the development
	// organization, all Scrum Teams must follow it as a minimum.
	// If "Done" for an increment is not a convention of the development organization, the
	// Development Team of the Scrum Team must define a definition of “Done” appropriate for the
	// product. If there are multiple Scrum Teams working on the system or product release, the
	// Development Teams on all the Scrum Teams must mutually define the definition of “Done.”
	// Each Increment is additive to all prior Increments and thoroughly tested, ensuring that all
	// Increments work together.
	// As Scrum Teams mature, it is expected that their definitions of “Done” will expand to include
	// more stringent criteria for higher quality. New definitions, as used, may uncover work to be
	// done in previously “Done” increments. Any one product or system should have a definition of
	// “Done” that is a standard for any work done on it.
	Done string
	Scope
	work string
	Planning
	DailyScrums []DailyScrum
	Events      []Event
	// The Product Backlog items selected for this Sprint plus the plan for delivering
	// them is called the Sprint Backlog.
	// 	The Sprint Backlog makes visible all the work that the Development Team identifies as necessary
	// to meet the Sprint Goal. To ensure continuous improvement, it includes at least one high
	// priority process improvement identified in the previous Retrospective meeting.
	//
	// 	The Sprint Backlog is a plan with enough detail that changes in progress can be understood in
	// the Daily Scrum. The Development Team modifies the Sprint Backlog throughout the Sprint, and
	// the Sprint Backlog emerges during the Sprint. This emergence occurs as the Development Team
	// works through the plan and learns more about the work needed to achieve the Sprint Goal.
	//
	// 	As new work is required, the Development Team adds it to the Sprint Backlog. As work is
	// performed or completed, the estimated remaining work is updated. When elements of the plan
	// are deemed unnecessary, they are removed. Only the Development Team can change its Sprint
	// Backlog during a Sprint. The Sprint Backlog is a highly visible, real-time picture of the work that
	// the Development Team plans to accomplish during the Sprint, and it belongs solely to the
	// Development Team.
	Backlogs []Backlog
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
