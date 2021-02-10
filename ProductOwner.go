package main

// ProductOwner struct
// The Product Owner is responsible for maximizing the value of the product resulting from work
// of the Development Team.
// The Product Owner is the sole person responsible for managing the Product Backlog. Product
// Backlog management includes:
// • Clearly expressing Product Backlog items;
// • Ordering the items in the Product Backlog to best achieve goals and missions;
// • Optimizing the value of the work the Development Team performs;
// • Ensuring that the Product Backlog is visible, transparent, and clear to all, and shows what
// the Scrum Team will work on next; and,
// • Ensuring the Development Team understands items in the Product Backlog to the level
// needed.
// The Product Owner may do the above work, or have the Development Team do it. However, the
// Product Owner remains accountable.
type ProductOwner struct {
}

// NewProductOwner creates a ProductOwner struct.
func NewProductOwner() *ProductOwner {
	return &ProductOwner{}
}

// Filter returns a filterd value.
func (p *ProductOwner) Filter(x float64) float64 {
	return x
}
