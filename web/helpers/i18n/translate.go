package i18n

import "github.com/mraron/njudge/utils/problems"

func TranslateContent(locale string, cs problems.Contents) problems.LocalizedData {
	search := func(loc string) (problems.LocalizedData, bool) {
		for _, c := range cs {
			if locale == c.Locale() {
				return c, true
			}
		}

		return problems.BytesData{}, false
	}

	if val, ok := search(locale); ok {
		return val
	}

	if val, ok := search("hungarian"); ok {
		return val
	}

	if len(cs) == 0 {
		return problems.BytesData{Loc: "-", Val: []byte("undefined"), Typ: "text"}
	}
	return cs[0]
}
