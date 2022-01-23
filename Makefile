.PHONY:run
run: build
	@./remote-bmi

.PHONY: build
build:
	@cd client && npm run build
	@go build -o remote-bmi server/*.go
