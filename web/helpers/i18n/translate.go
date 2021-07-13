package i18n

import "github.com/mraron/njudge/utils/problems"

func TranslateContent(locale string, cs problems.Contents) problems.Content {
	search := func(loc string) (problems.Content, bool) {
		for _, c := range cs {
			if locale == c.Locale {
				return c, true
			}
		}

		return problems.Content{}, false
	}

	if val, ok := search(locale); ok {
		return val
	}

	if val, ok := search("hungarian"); ok {
		return val
	}

	if len(cs) == 0 {
		return problems.Content{"-", []byte("undefined"), "text"}
	}
	return cs[0]
}