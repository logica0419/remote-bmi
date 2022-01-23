.PHONY:run
run:
	go run server/*.go

.PHONY: build
build:
	go build server/*.go -o remote-bmi
