####
## The Renty Backend Stack Makefile
## (run $ make for help)
####
.PHONY: help start

RM	     = rm -f
APP_NAME     := $(shell head -n 1 README.md | cut -d ' ' -f2 |  tr '[:upper:]' '[:lower:]')
APP_VSN      := $(shell git describe --tags | cut -d '-' -f1)
BUILD        := $(shell git rev-parse --short HEAD)

SHELL        := /bin/bash

STAGE        ?= dev
CREDENTIALS  = .credential
ENVFILE      = '.$(strip $(STAGE)).env'

EXTERNALIP   = $(shell curl -s ifconfig.co)


help: ## Print this help message and exit
	@echo -e "\n\t$(APP_NAME):$(APP_VSN)-$(BUILD) \033[1mmake\033[0m options:\n"
	@perl -nle 'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "- \033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo -e "\n"

default: help


dbseed: ## Seeding the database with ./backend/seed
	@cd backend/seed && go run .
	@echo "Et voila ! 🌾"

dbdrop: ## Easily empty the whole database content and user
	@cd backend/seed && go run . drop
	@echo "Et voila ! 🧹"

dbreset: ## Shortcut for drop and seed
	@make dbdrop
	@make dbseed
	@echo "Et voila ! 🚿"

scrapping: ## Run (mostly Leboncoin) scrapping
	@cd backend/scrappers/leboncoin && ./start.sh
	@echo "Et voila ! 🕷️ "

codegen: ## Let gqlgen generate resolver
	@cd backend/api && go run github.com/99designs/gqlgen .
	@echo "Et voila ! 🪄"

changeregistry: ## Change container registry (from treescale to gitlab or the contrary)
	@./dev-setup/scripts/change_registry.sh
	@echo "Et voila ! 🔧"

updategodep: ## Update goscrappy dep in every service
	@./dev-setup/scripts/update_godep.sh
	@echo "Et voila ! 🔧"

kubeinstall: ## Install local k3s cluster and set container registry secrets
	@exec ./dev-setup/scripts/k3s_install.sh

serviceinstall: ## Install our stack on your local k3s cluster
	@./dev-setup/scripts/services_install.sh

openfaassecrets: ## Upsert OpenFAAS functions secrets (create or update)
	@./infra/openfaas/set_registry_secrets.sh
	@./infra/openfaas/set_db_secrets.sh
	@./infra/openfaas/set_jwt_secrets.sh
	@./infra/openfaas/set_smtp_secrets.sh
