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

.PHONY: build-client
build-client:
	@cd client && npm run build

.PHONY: build
build: build-client
	@go build -o remote-bmi server/*.go

.PHONY: help
help: build
	@./remote-bmi -h

.PHONY: config
config: build
	@./remote-bmi config

.PHONY: serve
serve: build
	@./remote-bmi serve

.PHONY: dev
dev:
	@cd client && npm run dev
