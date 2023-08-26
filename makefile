.PHONY: all

GO_CMD=go
GO_TEST=$(GO_CMD) test
GO_TOOL=$(GO_CMD) tool
COVERAGE_PROFILE_NAME=coverage.out

test:
	$(GO_TEST) -v ./... -covermode=count -coverprofile=$(COVERAGE_PROFILE_NAME) -tags mock
	$(GO_TOOL) cover -func=coverage.out -o=coverage.out