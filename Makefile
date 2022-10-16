help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build local working copy
	@echo "Fetching packages"
	@go get -v
	@echo "Building binary"
	@go build -o vlc-pl-convert
