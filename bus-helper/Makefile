.PHONY: fmt
fmt: ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: build
build: ## Build a version
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./...

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: help
# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: upload
upload:
	scp -P 222 bus-helper root@47.52.232.228:/root

.DEFAULT_GOAL := help