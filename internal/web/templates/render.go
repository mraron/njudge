package templates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

const (
	CSRFTokenContextKey = "_csrf"
	CSRFTokenLookup     = "form:" + CSRFTokenContextKey

	UserContextKey    = "user"
	URLPathContextKey = "_url_path"
	TitleContextKey   = "_njudge_title"

	UsersContextKey         = "_njudge_users"
	ProblemsContextKey      = "_njudge_problems"
	ProblemsStoreContextKey = "_njudge_problems_store"
	PartialsStoreContextKey = "_njudge_partials_store"
)

func Middleware(users njudge.Users, ps njudge.Problems, problemStore problems.Store, partialsStore Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(URLPathContextKey, templ.SafeURL(c.Request().URL.Path))
			c.Set(UsersContextKey, users)
			c.Set(ProblemsContextKey, ps)
			c.Set(ProblemsStoreContextKey, problemStore)
			c.Set(PartialsStoreContextKey, partialsStore)
			return next(c)
		}
	}
}

func partial(ctx context.Context, name string) string {
	if store, ok := ctx.Value(PartialsStoreContextKey).(Store); ok {
		if res, err := store.Get(name); err == nil {
			return res
		}
	}
	return ""
}

func userContext(ctx context.Context) *njudge.User {
	if u, ok := ctx.Value("user").(*njudge.User); ok {
		return u
	}
	return nil
}

func user(ctx context.Context, id int) *njudge.User {
	if users, ok := ctx.Value(UsersContextKey).(njudge.Users); ok {
		if u, err := users.Get(ctx, id); err == nil {
			return u
		}
	}
	return nil
}

func problem(ctx context.Context, id int) *njudge.Problem {
	if ps, ok := ctx.Value(ProblemsContextKey).(njudge.Problems); ok {
		if p, err := ps.Get(ctx, id); err == nil {
			return p
		}
	}
	return nil
}

func problemWithStored(ctx context.Context, p *njudge.Problem) *njudge.ProblemStoredData {
	if p != nil {
		if store, ok := ctx.Value(ProblemsStoreContextKey).(problems.Store); ok {
			if res, err := p.WithStoredData(store); err == nil {
				return &res
			}
			return nil
		}
	}
	return nil
}

func d(x int) string {
	return strconv.Itoa(x)
}

func f(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func memKib(a memory.Amount) string {
	return d(int(a / memory.KiB))
}

func memMiB(a memory.Amount) string {
	return d(int(a / memory.MiB))
}

func iif[T any](b bool, i, e T) T {
	if b {
		return i
	}
	return e
}

func TrCs(ctx context.Context, cs problems.Contents) problems.LocalizedData {
	tr := ctx.Value(i18n.TranslatorContextKey).(i18n.Translator)
	return tr.TranslateContent(cs)
}

func Tr(ctx context.Context, key string, args ...any) string {
	tr := ctx.Value(i18n.TranslatorContextKey).(i18n.Translator)
	return tr.Translate(key, args...)
}

type echoContextWrapper struct {
	echo.Context
}

func (e echoContextWrapper) Deadline() (deadline time.Time, ok bool) {
	return e.Request().Context().Deadline()
}

func (e echoContextWrapper) Done() <-chan struct{} {
	return e.Request().Context().Done()
}

func (e echoContextWrapper) Err() error {
	return e.Request().Context().Err()
}

func (e echoContextWrapper) Value(key any) any {
	if _, ok := key.(string); !ok {
		return e.Request().Context().Value(key)
	}
	return e.Get(key.(string))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(echoContextWrapper{ctx}, buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
