module github.com/mraron/njudge

go 1.22

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2 // indirect
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/gorilla/sessions v1.2.2
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/labstack/echo-contrib v0.17.1
	github.com/labstack/echo/v4 v4.12.0
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/lib/pq v1.10.9
	github.com/markbates/goth v1.80.0
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/sendgrid/rest v2.6.9+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.14.0+incompatible
	github.com/spf13/afero v1.11.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/randomize v0.0.1 // indirect
	github.com/volatiletech/sqlboiler/v4 v4.16.2
	github.com/volatiletech/strmangle v0.0.6
	golang.org/x/crypto v0.31.0
	golang.org/x/oauth2 v0.20.0 // indirect
)

require (
	github.com/PuerkitoBio/goquery v1.9.2
	github.com/a-h/templ v0.2.778
	github.com/antonlindstrom/pgstore v0.0.0-20220421113606-e3a6e3fed12a
	github.com/erni27/imcache v1.2.0
	github.com/karrick/gobls v1.3.5
	github.com/quasoft/memstore v0.0.0-20191010062613-2bce066d2b0b
	github.com/samber/slog-echo v1.14.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.9.0
	github.com/yuin/goldmark v1.7.1
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
	golang.org/x/net v0.28.0
	golang.org/x/sync v0.10.0
	golang.org/x/text v0.21.0
	golang.org/x/time v0.5.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/ericlagergren/decimal v0.0.0-20190420051523-6335edbaa640 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

replace github.com/karrick/gobls v1.3.5 => github.com/mraron/gobls v0.0.0-20240125210852-e9cc54e22010

replace github.com/golang-migrate/migrate/v4 v4.17.1 => github.com/mraron/migrate/v4 v4.16.0
