DIR_LIST=go list -f '{{.Dir}}' ./... | grep -v /vendor/
PKG_LIST=go list ./... | grep -v /vendor/

.PHONY: check
check: format lint vet test-race

.PHONY: deps
deps:
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: format
format:
	gofmt -w `$(call DIR_LIST)`

.PHONY: lint
lint:
	golint -set_exit_status `${PKG_LIST}`

.PHONY: test
test:
	go test -a -mod=vendor `$(call PKG_LIST)`

.PHONY: test-race
test-race:
	go test -a -mod=vendor -race `$(call PKG_LIST)`

.PHONY: test-coverage
test-coverage:
	go test -a -mod=vendor -cover `$(call PKG_LIST)`

.PHONY: vet
vet:
	go vet -mod=vendor `$(call PKG_LIST)`