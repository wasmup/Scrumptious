package main

// Review struct
// A “Done” increment is required at the Sprint Review.
// Only members of the Development Team create the Increment
//
// A Sprint Review is held at the end of the Sprint to inspect the Increment and adapt the Product
// Backlog if needed. During the Sprint Review, the Scrum Team and stakeholders collaborate
// about what was done in the Sprint. Based on that and any changes to the Product Backlog
// during the Sprint, attendees collaborate on the next things that could be done to optimize value.
// This is an informal meeting, not a status meeting, and the presentation of the Increment is
// intended to elicit feedback and foster collaboration.
// This is at most a four-hour meeting for one-month Sprints. For shorter Sprints, the event is
// usually shorter. The Scrum Master ensures that the event takes place and that attendees
// understand its purpose. The Scrum Master teaches everyone involved to keep it within the timebox.
//
// The Sprint Review includes the following elements:
// • Attendees include the Scrum Team and key stakeholders invited by the Product Owner;
// • The Product Owner explains what Product Backlog items have been “Done” and what has
// not been “Done”;
// • The Development Team discusses what went well during the Sprint, what problems it ran
// into, and how those problems were solved;
// • The Development Team demonstrates the work that it has “Done” and answers questions
// about the Increment;
// • The Product Owner discusses the Product Backlog as it stands. He or she projects likely
// target and delivery dates based on progress to date (if needed);
// • The entire group collaborates on what to do next, so that the Sprint Review provides
// valuable input to subsequent Sprint Planning;
// • Review of how the marketplace or potential use of the product might have changed what is
// the most valuable thing to do next; and,
// • Review of the timeline, budget, potential capabilities, and marketplace for the next
// anticipated releases of functionality or capability of the product
type Review struct {
	Progress
	Sprint
	// 	The result of the Sprint Review is a revised Product Backlog that defines the probable Product
	// Backlog items for the next Sprint. The Product Backlog may also be adjusted overall to meet new
	// opportunities
	Next Sprint
}

// NewReview creates a Review struct.
func NewReview() *Review {
	return &Review{}
}

// Filter returns a filterd value.
func (p *Review) Filter(x float64) float64 {
	return x
}
