package main

// ScrumMaster struct
// Scrum Master Service to the Product Owner
// The Scrum Master serves the Product Owner in several ways, including:
// • Ensuring that goals, scope, and product domain are understood by everyone on the Scrum
// Team as well as possible;
// • Finding techniques for effective Product Backlog management;
// • Helping the Scrum Team understand the need for clear and concise Product Backlog items;
// • Understanding product planning in an empirical environment;
// • Ensuring the Product Owner knows how to arrange the Product Backlog to maximize value;
// • Understanding and practicing agility; and,
// • Facilitating Scrum events as requested or needed.
// Scrum Master Service to the Development Team
// The Scrum Master serves the Development Team in several ways, including:
// • Coaching the Development Team in self-organization and cross-functionality;
// • Helping the Development Team to create high-value products;
// • Removing impediments to the Development Team’s progress;
// • Facilitating Scrum events as requested or needed; and,
// • Coaching the Development Team in organizational environments in which Scrum is not yet
// fully adopted and understood.
// Scrum Master Service to the Organization
// The Scrum Master serves the organization in several ways, including:
// • Leading and coaching the organization in its Scrum adoption;
// • Planning Scrum implementations within the organization;
// • Helping employees and stakeholders understand and enact Scrum and empirical product
// development;
// • Causing change that increases the productivity of the Scrum Team; and,
// • Working with other Scrum Masters to increase the effectiveness of the application of Scrum
// in the organization
//
// The Scrum Master ensures that the Development Team has the meeting, but the Development
// Team is responsible for conducting the Daily Scrum. The Scrum Master teaches the
// Development Team to keep the Daily Scrum within the 15-minute time-box.
// The Daily Scrum is an internal meeting for the Development Team. If others are present, the
// Scrum Master ensures that they do not disrupt the meeting.
type ScrumMaster struct {
}

// NewScrumMaster creates a ScrumMaster struct.
func NewScrumMaster() *ScrumMaster {
	return &ScrumMaster{}
}

// Filter returns a filterd value.
func (p *ScrumMaster) Filter(x float64) float64 {
	return x
}
