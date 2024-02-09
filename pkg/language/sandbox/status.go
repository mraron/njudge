package sandbox

import "time"

type Verdict int

const (
	VerdictOK Verdict = 1 << iota
	VerdictTL
	VerdictML
	VerdictRE
	VerdictXX
	VerdictCE
)

func (v Verdict) String() string {
	switch v {
	case VerdictOK:
		return "OK"
	case VerdictTL:
		return "TL"
	case VerdictML:
		return "ML"
	case VerdictRE:
		return "RE"
	case VerdictXX:
		return "XX"
	case VerdictCE:
		return "CE"
	}
	return "??"
}

type Status struct {
	Verdict  Verdict
	Signal   int
	Memory   int
	Time     time.Duration
	ExitCode int
}
