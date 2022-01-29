.PHONY: init
init:
	go mod download && cd client && npm i

.PHONY: wire
wire:
	@cd server/cmd && wire

.PHONY: tbls
tbls:
	@rm -rf server/docs/db_schema
	@cd server/docs && tbls doc

.PHONY: build
build: build-server build-client

.PHONY: build-client
build-client:
	@cd client && npm run build

.PHONY: build-server
build-server:
	@go build -o remote-bmi server/*.go

.PHONY: help
help: build-server
	@./remote-bmi -h

.PHONY: config
config: build-server
	@./remote-bmi config

.PHONY: serve
serve: build-server
	@./remote-bmi serve

.PHONY: dev
dev:
	@cd client && npm run dev
