package main

// DailyScrum struct
type DailyScrum struct {
}

// NewDailyScrum creates a DailyScrum struct.
func NewDailyScrum() *DailyScrum {
	return &DailyScrum{}
}

// Filter returns a filterd value.
func (p *DailyScrum) Filter(x float64) float64 {
	return x
}
