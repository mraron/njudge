package i18n

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net/http"
	"time"

	_ "github.com/mraron/njudge/internal/web/translations"
)

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

type Translator struct {
	ID         string
	LocaleName string

	printer *message.Printer
}

var locales = []Translator{
	{
		ID:         "hu-HU",
		LocaleName: "hungarian",
		printer:    message.NewPrinter(language.MustParse("hu-HU")),
	},
	{
		ID:         "en-US",
		LocaleName: "english",
		printer:    message.NewPrinter(language.MustParse("en-US")),
	},
}

func Get(id string) (Translator, bool) {
	for _, locale := range locales {
		if locale.ID == id {
			return locale, true
		}
	}

	return Translator{}, false
}

func (t Translator) Translate(key message.Reference, args ...interface{}) string {
	return t.printer.Sprintf(key, args...)
}

func (t Translator) TranslateContent(cs problems.Contents) problems.LocalizedData {
	return TranslateContent(t.LocaleName, cs)
}

func SetTranslatorMiddleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			trySetting := func(langCode string) bool {
				if t, ok := Get(langCode); ok {
					c.Set("translator", t)
					c.SetCookie(&http.Cookie{
						Name:    "lang",
						Value:   t.ID,
						Path:    "/",
						Expires: time.Now().Add(24 * time.Hour),
					})
					return true
				}

				return false
			}

			if trySetting(c.QueryParam("lang")) {
				return next(c)
			}

			cookie, err := c.Cookie("lang")
			if err == nil {
				if trySetting(cookie.Value) {
					return next(c)
				}
			}

			if trySetting("hu-HU") {
				return next(c)
			}
			return errors.New("can't set language")
		}
	}
}
