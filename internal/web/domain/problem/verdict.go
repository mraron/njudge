package problem

import (
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/problems"
)

type VerdictType int

const (
	VerdictTypeRunning VerdictType = iota
	VerdictTypeWrong
	VerdictTypeAC
)

type Verdict int

const (
	VerdictAC = Verdict(problems.VerdictAC)
	VerdictWA = Verdict(problems.VerdictWA)
	VerdictRE = Verdict(problems.VerdictRE)
	VerdictTL = Verdict(problems.VerdictTL)
	VerdictML = Verdict(problems.VerdictML)
	VerdictXX = Verdict(problems.VerdictXX)
	VerdictDR = Verdict(problems.VerdictDR)
	VerdictPC = Verdict(problems.VerdictPC)
	VerdictPE = Verdict(problems.VerdictPE)

	VerdictCE Verdict = 998
	VerdictRU Verdict = 999
	VerdictUP Verdict = 1000
)

func VerdictFromProblemsVerdictName(name problems.VerdictName) Verdict {
	return Verdict(name)
}

func (v Verdict) String() string {
	switch v {
	case VerdictAC:
		return "Accepted"
	case VerdictWA:
		return "Wrong answer"
	case VerdictRE:
		return "Runtime error"
	case VerdictTL:
		return "Time limit exceeded"
	case VerdictML:
		return "Memory limit exceeded"
	case VerdictXX:
		return "Internal error"
	case VerdictDR:
		return "Didn't run"
	case VerdictPC:
		return "Partially correct"
	case VerdictPE:
		return "Presentation error"
	case VerdictCE:
		return "Compilation error"
	case VerdictRU:
		return "Running"
	case VerdictUP:
		return "Uploaded"
	}

	return ""
}

func (v Verdict) VerdictType() VerdictType {
	if v == VerdictRU || v == VerdictUP {
		return VerdictTypeRunning
	}
	if v == VerdictAC {
		return VerdictTypeAC
	}

	return VerdictTypeWrong
}

func (v Verdict) Translate(t i18n.Translator) string {
	switch v {
	case VerdictAC:
		return t.Translate("Accepted")
	case VerdictWA:
		return t.Translate("Wrong answer")
	case VerdictRE:
		return t.Translate("Runtime error")
	case VerdictTL:
		return t.Translate("Time limit exceeded")
	case VerdictML:
		return t.Translate("Memory limit exceeded")
	case VerdictXX:
		return t.Translate("Internal error")
	case VerdictDR:
		return t.Translate("Didn't run")
	case VerdictPC:
		return t.Translate("Partially correct")
	case VerdictPE:
		return t.Translate("Presentation error")
	case VerdictCE:
		return t.Translate("Compilation error")
	case VerdictRU:
		return t.Translate("Running")
	case VerdictUP:
		return t.Translate("Uploaded")
	}

	return ""
}
