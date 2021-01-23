DIR_LIST=go list -f '{{.Dir}}' ./... | grep -v /vendor/
PKG_LIST=go list ./... | grep -v /vendor/

.PHONY: ci
ci:
	golangci-lint run

.PHONY: deps
deps:
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: format
format:
	gofmt -w `$(call DIR_LIST)`

.PHONY: test
test:
	go test -mod=vendor -short `$(call PKG_LIST)`

.PHONY: test-coverage
test-coverage:
	go test -mod=vendor -race -cover -coverprofile=coverage.out `$(call PKG_LIST)`