.PHONY: wire
wire:
	@cd server/cmd && wire

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

.PHONY:config
config: build-server
	@./remote-bmi config

.PHONY:serve
serve: build
	@./remote-bmi serve

.PHONY:dev
dev:
	@cd client && npm run dev
