package njudge

type Judge struct {
	ID           int
	Name         string
	URL          string
	Online       bool
	ProblemList  []string
	LanguageList []string
}
