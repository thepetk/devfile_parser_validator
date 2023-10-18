.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: gomod_tidy
gomod_tidy: ## Update Go mod dependencies
	 go mod tidy

.PHONY: gofmt
gofmt: ## Run Go formatting checks
	go fmt -x ./...

.PHONY: build
build: ## Build executable without extension (Unix)
	go build -o devfile_parser

.PHONY: buildWin
buildWin: ## Build executable with Windows .exe extension
	go build -o devfile_parser.exe