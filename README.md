# njudge

[![Tests](https://github.com/mraron/njudge/actions/workflows/tests.yml/badge.svg)](https://github.com/mraron/njudge/actions/workflows/tests.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mraron/njudge.svg)](https://pkg.go.dev/github.com/mraron/njudge)
[![Go Report Card](https://goreportcard.com/badge/github.com/mraron/njudge)](https://goreportcard.com/report/github.com/mraron/njudge)

Online Judge system written in golang. A running version can be found [here](https://njudge.hu).

It consists of two main parts:
* A group of reusable components, located in `pkg/`, which can be used in a variety ways related to the conversion, running and modification of algorithmic/olympiad tasks.
* A web based online judge system, located in `internal/`. 

<!--
## Features
@TODO
-->

## How to run the online judge

Disclaimer: it's possible to run the system without docker, but it's much more of a hassle to set up the environment. 

First, build the containers:

```
docker-compose build base
docker-compose build
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
docker-compose up
```

<!--@TODO create admin account -->

To make problems available to the system you have two options: 
1. Use the automatically created `njudge_problems` volume, copy the problems there via `docker cp`.
2. Modify the `docker-compose.yml` to use a bind mount pointing to a directory on your host system. 

If both the judge and web component sees the problems, you have to use the admin panel (or some sql client) to create the necessary `ProblemRel`s for it to be displayed on the frontend.
