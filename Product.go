package main

// Product struct
type Product struct {
	Backlog []Backlog
	Sprints []Sprint
	ProductOwner
	accepted bool
	// estimated remaining work
}

// NewProduct creates a Product struct.
func NewProduct() *Product {
	return &Product{}
}

// Filter returns a filterd value.
func (p *Product) Filter(x float64) float64 {
	return x
}
