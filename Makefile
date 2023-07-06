SHELL=/bin/bash
.DEFAULT_GOAL=setup
CURRENTDIR=$(shell dirname `pwd`)

ifneq (,$(wildcard ./.env))
include .env
export
endif

setup: go.mod
	@echo "`tput bold`#### Verifying configuration files and server certificates ####`tput sgr0`"
	@test -f cert.pem || go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
	@test -f .env || cp .env.example .env
	@test -f config.yaml || cp config.example.yaml config.yaml
	@make generate_key
	@echo "## Configuration files are now ready to use ##"

	@sleep 2

	@echo "`tput bold`#### Installing dependencies to your project ####`tput sgr0`"
	go mod tidy

	go get -u golang.org/x/lint/golint
	go install golang.org/x/lint/golint
	go get -u github.com/mgechev/revive
	go install github.com/mgechev/revive

	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	go install mvdan.cc/gofumpt@latest
	@sleep 2

	@echo "creating .venv and installing it's dependencies"
	test -d .venv || python3 -m venv .venv
	. .venv/bin/activate; pip install pymigratedb

	@echo "## All dependencies installed successfully ##"
	@sleep 2

	@echo ""
	@echo "`tput bold``tput setaf 1`## Verify config.yaml and .env and fill it according to your params ##`tput sgr0`"
	@echo ""

run: .env
	TEST_DATABASE="" go run .

generate_key:
	test -f private.pem || openssl genpkey -out private.pem -algorithm RSA -pkeyopt rsa_keygen_bits:4096

generate_public:
	test -f private.pem && openssl pkey -in private.pem -pubout -out public.pem

migrate: .venv
	. .venv/bin/activate; migrate $(command) --driver pgsql

create_migration: .venv
	. .venv/bin/activate; migrate create --driver pgsql --migration_name $(name)

lint:
  ifndef file
	$(error file is not defined)
  else
	golangci-lint run $(file) --go=1.20 --enable-all --disable tagliatelle,wsl,godox,lll,gochecknoglobals,exhaustruct,exhaustivestruct,wrapcheck
  endif

TEST_DATABASE=test_db
test:
	. .venv/bin/activate; migrate execute --driver pgsql
	go clean -testcache;
	go test -tags=unit -short -timeout 30s -v ./...

format:
  ifndef file
	$(error file is not defined)
  else
	gofumpt -w $(file)
  endif

help:
	@echo "List of Makefile commands"