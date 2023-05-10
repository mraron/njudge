module github.com/mraron/njudge

go 1.18

require (
	cloud.google.com/go/compute v1.5.0 // indirect
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/labstack/echo-contrib v0.12.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/labstack/gommon v0.3.1
	github.com/lib/pq v1.10.5
	github.com/markbates/goth v1.69.0
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/sendgrid/rest v2.6.8+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.11.0+incompatible
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/spf13/afero v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/randomize v0.0.1 // indirect
	github.com/volatiletech/sqlboiler/v4 v4.13.0
	github.com/volatiletech/strmangle v0.0.4
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.uber.org/multierr v1.8.0
	golang.org/x/crypto v0.1.0
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/antonlindstrom/pgstore v0.0.0-20220421113606-e3a6e3fed12a
	github.com/gomarkdown/markdown v0.0.0-20220905174103-7b278df48cfb
	github.com/google/go-cmp v0.5.8
	github.com/karrick/gobls v1.3.5
	go.uber.org/zap v1.23.0
	golang.org/x/exp v0.0.0-20220827204233-334a2380cb91
	golang.org/x/text v0.3.8
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
)

replace github.com/golang-migrate/migrate/v4 => github.com/mraron/migrate/v4 v4.15.2-0.20220427062446-a7e44b82b3fd
