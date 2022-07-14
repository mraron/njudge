# njudge

[![Tests](https://github.com/mraron/njudge/actions/workflows/tests.yml/badge.svg)](https://github.com/mraron/njudge/actions/workflows/tests.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/mraron/njudge)](https://goreportcard.com/report/github.com/mraron/njudge)

Online Judge system written in golang.

## Installation

Build and start the containers:

```
docker-compose build
docker-compose up
```

Run migrations:
```
docker exec -it judge_web_1 ./njudge web migrate --up
```

Modify the environment files `web.env` and `judge.env` (restart the containers for the changes to take effect).

Add problems to the problems volume, both the web and judge will detect them. 

## Modules
* njudge/web: web frontend [![](https://godoc.org/github.com/mraron/njudge/web?status.svg)](http://godoc.org/github.com/mraron/njudge/web)
* njudge/judge: judges programs sent to it by an njudge/web instance [![](https://godoc.org/github.com/mraron/njudge/judge?status.svg)](http://godoc.org/github.com/mraron/njudge/judge)
* njudge/utils/language: sandboxing utilities [![](https://godoc.org/github.com/mraron/njudge/utils/language?status.svg)](http://godoc.org/github.com/mraron/njudge/utils/language)
* njudge/utils/problems: parsing and internal representation of problems [![](https://godoc.org/github.com/mraron/njudge/utils/problems?status.svg)](https://godoc.org/github.com/mraron/njudge/utils/problems)

## Screenshots
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/1.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/2.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/3.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/4.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/5.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/6.png" width="25%" height="25%">

