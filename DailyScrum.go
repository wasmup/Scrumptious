package main

// DailyScrum struct
// The Daily Scrum is a 15-minute time-boxed event for the Development Team. The Daily Scrum is
// held every day of the Sprint. At it, the Development Team plans work for the next 24 hours. This
// optimizes team collaboration and performance by inspecting the work since the last Daily Scrum
// and forecasting upcoming Sprint work. The Daily Scrum is held at the same time and place each
// day to reduce complexity.
//
// The Development Team uses the Daily Scrum to inspect progress toward the Sprint Goal and to
// inspect how progress is trending toward completing the work in the Sprint Backlog. The Daily
// Scrum optimizes the probability that the Development Team will meet the Sprint Goal. Every
// day, the Development Team should understand how it intends to work together as a selforganizing team to accomplish the Sprint Goal and create the anticipated Increment by the end
// of the Sprint.
//
// The structure of the meeting is set by the Development Team and can be conducted in different
// ways if it focuses on progress toward the Sprint Goal. Some Development Teams will use
// questions, some will be more discussion based. Here is an example of what might be used:
// • What did I do yesterday that helped the Development Team meet the Sprint Goal?
// • What will I do today to help the Development Team meet the Sprint Goal?
// • Do I see any impediment that prevents me or the Development Team from meeting the
// Sprint Goal?
// The Development Team or team members often meet immediately after the Daily Scrum for
// detailed discussions, or to adapt, or replan, the rest of the Sprint’s work.
//
// Daily Scrums improve communications, eliminate other meetings, identify impediments to
// development for removal, highlight and promote quick decision-making, and improve the
// Development Team’s level of knowledge. This is a key inspect and adapt meeting.
type DailyScrum struct {
	impediments []Impediment
}

// NewDailyScrum creates a DailyScrum struct.
func NewDailyScrum() *DailyScrum {
	return &DailyScrum{}
}

// Filter returns a filterd value.
func (p *DailyScrum) Filter(x float64) float64 {
	return x
}
