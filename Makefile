test:
	go test

format:
	goimports -w errorhandle.go

build:
	cd ./cmd/errorhandle && go build -o errorhandle

.PHONY: test format build

