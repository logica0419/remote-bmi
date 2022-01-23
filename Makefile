.PHONY: wire
wire:
	@cd server/cmd && wire

.PHONY: build
build:
	@cd client && npm run build
	@go build -o remote-bmi server/*.go

.PHONY:run
run: build
	@./remote-bmi
