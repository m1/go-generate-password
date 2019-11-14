PACKAGES := $(shell go list ./... | grep -v cmd)

.PHONY: test
test:
	go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out $(PACKAGES)