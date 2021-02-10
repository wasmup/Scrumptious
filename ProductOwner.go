package main

// ProductOwner struct
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
