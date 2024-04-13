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

up: build ## builds an runs docker-compose up
	COMPOSE_PROJECT_NAME="$(PROJECT_NAME)" docker-compose up

