# docker Makefile based on https://gist.github.com/mpneuried/0594963ad38e68917ef189b4e6a269db

PROJECT_NAME=njudge

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: build_base build_web build_glue build_judge ## Builds all services

build_base: ## builds njudge-base
	docker build -t $(PROJECT_NAME)-base .
build_web:  ## builds njudge-web
	docker build --build-arg="PROJECT_NAME=$(PROJECT_NAME)" -t $(PROJECT_NAME)-web -f web.Dockerfile .
build_glue: ## builds njudge-glue
	docker build --build-arg="PROJECT_NAME=$(PROJECT_NAME)" -t $(PROJECT_NAME)-glue -f glue.Dockerfile .
build_judge:  ## builds njudge-judge
	docker build --build-arg="PROJECT_NAME=$(PROJECT_NAME)" -t $(PROJECT_NAME)-judge -f judge.Dockerfile .

gulp: ## run gulp
	npx gulp
lint: ## run golangci-lint linter
	golangci-lint run
generate: ## updates out.gotext.json translation files from source code (runs go generate)
	go generate ./...
translations: ## copies the out.gotext.json to the messages.gotext.json
	cp internal/web/translations/locales/en-US/out.gotext.json internal/web/translations/locales/en-US/messages.gotext.json
	cp internal/web/translations/locales/hu-HU/out.gotext.json internal/web/translations/locales/hu-HU/messages.gotext.json
models: ## updates internal/njudge/db/models from sqlboiler.toml
	sqlboiler psql
test: ## run tests
	go test ./...
templ:
	templ generate

up: build ## builds and runs docker-compose up
	COMPOSE_PROJECT_NAME="$(PROJECT_NAME)" docker-compose up

