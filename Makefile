test: format
	go test

format:
	goimports -w errorhandle.go

.PHONY: test format
