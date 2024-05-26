package njudge

type SolvedStatus string

const (
	Unattempted     SolvedStatus = "Unattempted"
	Attempted       SolvedStatus = "Attempted"
	PartiallySolved SolvedStatus = "PartiallySolved"
	Solved          SolvedStatus = "Solved"
	Unknown         SolvedStatus = "Unknown"
)
