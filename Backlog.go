package main

import "time"

// Backlog struct
// an accumulation of uncompleted work or matters needing to be dealt with.
// "the company took on extra staff to clear the backlog of work"
//
// The Product Backlog items selected for this Sprint plus the plan for delivering
// them is called the Sprint Backlog.
//
// Product Backlog
// The Product Backlog is an ordered list of everything that is known to be needed in the product.
// It is the single source of requirements for any changes to be made to the product. The Product
// Owner is responsible for the Product Backlog, including its content, availability, and ordering.
// A Product Backlog is never complete. The earliest development of it lays out the initially known
// and best-understood requirements. The Product Backlog evolves as the product and the
// environment in which it will be used evolves. The Product Backlog is dynamic; it constantly
// changes to identify what the product needs to be appropriate, competitive, and useful. If a
// product exists, its Product Backlog also exists.
// The Product Backlog lists all features, functions, requirements, enhancements, and fixes that
// constitute the changes to be made to the product in future releases. Product Backlog items have
// the attributes of a description, order, estimate, and value. Product Backlog items often include
// test descriptions that will prove its completeness when “Done.”
// As a product is used and gains value, and the marketplace provides feedback, the Product
// Backlog becomes a larger and more exhaustive list. Requirements never stop changing, so a
// Product Backlog is a living artifact. Changes in business requirements, market conditions, or
// technology may cause changes in the Product Backlog.
// Multiple Scrum Teams often work together on the same product. One Product Backlog is used
// to describe the upcoming work on the product. A Product Backlog attribute that groups items
// may then be employed.
//
// Product Backlog item
// During Product Backlog refinement,
// items are reviewed and revised. The Scrum Team decides how and when refinement is done.
// Refinement usually consumes no more than 10% of the capacity of the Development Team.
// However, Product Backlog items can be updated at any time by the Product Owner or at the
// Product Owner’s discretion.
//
// Higher ordered Product Backlog items are usually clearer and more detailed than lower ordered
// ones. More precise estimates are made based on the greater clarity and increased detail; the
// lower the order, the less detail. Product Backlog items that will occupy the Development Team
// for the upcoming Sprint are refined so that any one item can reasonably be “Done” within the
// Sprint time-box. Product Backlog items that can be “Done” by the Development Team within
// one Sprint are deemed “Ready” for selection in a Sprint Planning. Product Backlog items usually
// acquire this degree of transparency through the above described refining activities.
//
// The Development Team is responsible for all estimates. The Product Owner may influence the
// Development Team by helping it understand and select trade-offs, but the people who will
// perform the work make the final estimate.
//
// The Sprint Backlog is the set of Product Backlog items selected for the Sprint, plus a plan for
// delivering the product Increment and realizing the Sprint Goal. The Sprint Backlog is a forecast
// by the Development Team about what functionality will be in the next Increment and the work
// needed to deliver that functionality into a “Done” Increment.
type Backlog struct {
	Estimates time.Duration
	// 	Monitoring Sprint Progress
	// At any point in time in a Sprint, the total work remaining in the Sprint Backlog can be summed.
	// The Development Team tracks this total work remaining at least for every Daily Scrum to project
	// the likelihood of achieving the Sprint Goal. By tracking the remaining work throughout the
	// Sprint, the Development Team can manage its progress
	RemainingWork int
	Ready         bool
	Items         []string
	priority      int
}

// NewBacklog creates a Backlog struct.
func NewBacklog() *Backlog {
	return &Backlog{}
}

// Filter returns a filterd value.
func (p *Backlog) Filter(x float64) float64 {
	return x
}
