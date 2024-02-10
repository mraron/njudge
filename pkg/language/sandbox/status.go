package sandbox

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"time"
)

// Verdict is the overall outcome of running a program.
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

// Status contains information about how a program ran.
type Status struct {
	Verdict  Verdict
	Signal   int
	Time     time.Duration
	Memory   memory.Amount
	ExitCode int
}
