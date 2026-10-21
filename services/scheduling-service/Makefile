MIGRATIONS := migrations
TOOLS_BIN := $(CURDIR)/bin
GOOSE := $(TOOLS_BIN)/goose

ifneq (,$(wildcard .env))
include .env
export
endif

ifndef DATABASE_URL
$(error DATABASE_URL is not set. Copy .env.example to .env first.)
endif

.PHONY: deps migrate-up migrate-down migrate-status test fmt

deps:
	go mod download
	mkdir -p $(TOOLS_BIN)
	GOBIN=$(TOOLS_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1

migrate-up:
	$(GOOSE) -dir $(MIGRATIONS) postgres "$(DATABASE_URL)" up

migrate-down:
	$(GOOSE) -dir $(MIGRATIONS) postgres "$(DATABASE_URL)" down

migrate-status:
	$(GOOSE) -dir $(MIGRATIONS) postgres "$(DATABASE_URL)" status

test:
	go test ./...

fmt:
	gofmt -w .
