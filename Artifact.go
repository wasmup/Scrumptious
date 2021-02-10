package main

// Artifact struct
// Scrum Artifacts
// Scrum’s artifacts represent work or value to provide transparency and opportunities for
// inspection and adaptation. Artifacts defined by Scrum are specifically designed to maximize
// transparency of key information so that everybody has the same understanding of the artifact.
//
// Artifact Transparency
// Scrum relies on transparency. Decisions to optimize value and control risk are made based on
// the perceived state of the artifacts. To the extent that transparency is complete, these decisions
// have a sound basis. To the extent that the artifacts are incompletely transparent, these
// decisions can be flawed, value may diminish and risk may increase.
// The Scrum Master must work with the Product Owner, Development Team, and other involved
// parties to understand if the artifacts are completely transparent. There are practices for coping
// with incomplete transparency; the Scrum Master must help everyone apply the most
// appropriate practices in the absence of complete transparency. A Scrum Master can detect
// incomplete transparency by inspecting the artifacts, sensing patterns, listening closely to what is
// being said, and detecting differences between expected and real results.
// The Scrum Master’s job is to work with the Scrum Team and the organization to increase the
// transparency of the artifacts. This work usually involves learning, convincing, and change.
// Transparency doesn’t occur overnight, but is a path.
type Artifact struct {
}

// NewArtifact creates a Artifact struct.
func NewArtifact() *Artifact {
	return &Artifact{}
}

// Filter returns a filterd value.
func (p *Artifact) Filter(x float64) float64 {
	return x
}
