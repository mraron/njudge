package db

import "github.com/mraron/njudge/internal/njudge"

// TODO do a migration to strings in db
func NjudgeVerdictToDatabase(verdict njudge.Verdict) int {
	switch verdict {
	case njudge.VerdictAC:
		return 0
	case njudge.VerdictWA:
		return 1
	case njudge.VerdictRE:
		return 2
	case njudge.VerdictTL:
		return 3
	case njudge.VerdictML:
		return 4
	case njudge.VerdictXX:
		return 5
	case njudge.VerdictDR:
		return 6
	case njudge.VerdictPC:
		return 7
	case njudge.VerdictPE:
		return 8
	case njudge.VerdictCE:
		return 998
	case njudge.VerdictRU:
		return 999
	case njudge.VerdictUP:
		return 1000
	}
	return -1
}

func DatabaseVerdictToNjudge(verdict int) njudge.Verdict {
	switch verdict {
	case 0:
		return njudge.VerdictAC
	case 1:
		return njudge.VerdictWA
	case 2:
		return njudge.VerdictRE
	case 3:
		return njudge.VerdictTL
	case 4:
		return njudge.VerdictML
	case 5:
		return njudge.VerdictXX
	case 6:
		return njudge.VerdictDR
	case 7:
		return njudge.VerdictPC
	case 8:
		return njudge.VerdictPE
	case 998:
		return njudge.VerdictCE
	case 999:
		return njudge.VerdictRU
	case 1000:
		return njudge.VerdictUP
	}
	return njudge.VerdictUnknown
}
