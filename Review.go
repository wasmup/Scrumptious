package main

// Review struct
type Review struct {
}

// NewReview creates a Review struct.
func NewReview() *Review {
	return &Review{}
}

// Filter returns a filterd value.
func (p *Review) Filter(x float64) float64 {
	return x
}
