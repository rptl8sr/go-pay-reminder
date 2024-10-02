PROJECT := "go-pay-reminder"
V := "0.0.1"
USER := rptl8sr
EMAIL := $(USER)@gmail.com
LOCAL_BIN :=$(CURDIR)/bin
MIGRATIONS_DIR := migrations


.PHONY: git-init
git-init:
	gh repo create $(PROJECT) --private
	git init
	git config user.name "$(USER)"
	git config user.email "$(EMAIL)"
	git add Makefile go.mod
	git commit -m "Init commit"
	git remote add origin git@github.com:$(USER)/$(PROJECT).git
	git remote -v
	git push -u origin master


BN ?= dev
# make git-checkout BN=dev
.PHONY: git-checkout
git-checkout:
	git checkout -b $(BN)


.PHONY: blueprint
blueprint:
	touch .env
	echo '/.idea/\n/bin/\n\n*.env' > .gitignore
	mkdir -p config && touch config/config.yaml
	mkdir -p cmd && echo 'package main\n\nfunc main() {\n}\n' > cmd/main.go
	mkdir -p bin
	mkdir -p app && echo 'package app' > app/app.go
	mkdir -p internal/sheets && echo 'package sheets' > internal/sheets/sheets.go
	mkdir -p internal/config && echo 'package config' > internal/config/config.go
	mkdir -p internal/controller && echo 'package controller' > internal/controller/controller.go
	mkdir -p internal/logger && echo 'package logger' > internal/logger/logger.go
	mkdir -p internal/telegram && echo 'package telegram' >  internal/telegram/telegram.go


.PHONY: golangci-lint-install
golangci-lint-install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1


.PHONY: lint
lint:
	$(LOCAL_BIN)/golangci-lint run ./...


.PHONY: test
test:
	go vet ./...
	go test ./...


.PHONY: yzip
yzip:
	zip -r $(PROJECT)_`date +%Y-%m-%d`.zip configs internal go.mod go.sum handler.go credentials.json serviceCreds.json