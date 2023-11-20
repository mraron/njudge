package njudge

type SolvedStatus int

const (
	Unattempted SolvedStatus = iota
	Attempted
	PartiallySolved
	Solved
	Unknown
)
