# njudge

[![Tests](https://github.com/mraron/njudge/actions/workflows/tests.yml/badge.svg)](https://github.com/mraron/njudge/actions/workflows/tests.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mraron/njudge.svg)](https://pkg.go.dev/github.com/mraron/njudge)
[![Go Report Card](https://goreportcard.com/badge/github.com/mraron/njudge)](https://goreportcard.com/report/github.com/mraron/njudge)
[![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?style=for-the-badge&logo=discord&logoColor=white)](https://discord.gg/YYQSeeUweY)

Online Judge system written in golang. A running version can be found [here](https://njudge.hu).

It consists of two main parts:
* A group of reusable components, located in `pkg/`, which can be used in a variety ways related to the conversion, running and modification of algorithmic/olympiad tasks.
* A web based online judge system, located in `internal/`. 


## How to run the online judge

### Via docker
 
Build the containers:
```
make build
```

Then, run the migrations:
```
docker-compose run web ./njudge web migrate --up
```

Modify the environment files `web.env` and `judge.env`. Typically only `web.env` needs to be changed. A quite minimal `web.env` file for example would be:
```env
NJUDGE_MODE="production"
NJUDGE_HOSTNAME="yourhost.domain"
NJUDGE_URL="https://yourhost.domain"
NJUDGE_COOKIESECRET="somesecretstring"
```
Probably you would want some of `NJUDGE_SENDGRID`, `NJUDGE_SMTP`, `NJUDGE_GOOGLEAUTH` as well to provide OAuth login and email sending capabilities. 
For the full list of the configuration options supported by the web component check out [config.go](https://github.com/mraron/njudge/blob/v1-refactor/internal/web/helpers/config/config.go).

Then you can start all containers necessary via:
```
make up
```

<!--@TODO create admin account -->

To make problems available to the system you have two options: 
1. Use the automatically created `njudge_problems` volume, copy the problems there via `docker cp`.
2. Modify the `docker-compose.yml` to use a bind mount pointing to a directory on your host system. 

If both the judge and web component sees the problems, you have to use the admin panel (or some sql client) to create the necessary `ProblemRel`s for it to be displayed on the frontend.

### Vanilla method

Build everything with `go build`.  Start the services by executing all 3 commands: 
```
./njudge web
./njudge glue
./njudge judge
```

You most certainly need to configure a thing or two via `glue.yaml`, `judge.yaml` and `web.yaml`.

Example of `web.yaml`:
```yaml
problems_dir: /PATH_TO_PROBLEMS
db:
  user: your_user
  password: your_users_password
  host: localhost
  port: 1234567
  name: your_database
```
Example of `judge.yaml`:
```yaml
problems_dir: /PATH_TO_PROBLEMS
```
Example of `glue.yaml`:
```yaml
db:
  user: your_user
  password: your_users_password
  host: localhost
  port: 1234567
  name: your_database
```
For configuration check out [internal/web/config.go](internal/web/config.go) and [cmd/](cmd/) directory.
All values can be configured via command line arguments and environment variables as well. 

You will also need to execute migrations and add admin user. 

### Demo mode
WIP. It's a mode to run njudge without a database.

## Development

Use make to run and/or generate stuff.
* `make gulp` generates css and js (required `npm install`)
* `make lint` runs golangci-lint (requires [https://github.com/golangci/golangci-lint](golangci-lint) installed)
* `make translations` generates translation files
* `make models` generates models (requires [https://github.com/volatiletech/sqlboiler](sqlboiler) installed)
* `make test` runs tests
* `make templ` generates templates (requires [https://github.com/a-h/templ/](templ) installed)