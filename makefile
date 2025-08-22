help:
	@printf "Available targets:\n"
	@grep -E '^[1-9a-zA-Z_-]+:.*?## .*$$|(^#--)' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m %-43s\033[0m %s\n", $$1, $$2}' \
		| sed -e 's/\[32m #-- /[33m/'

default: help

#-- GO
build: ## Build Go binary
	go build -o bin/config .

intall: build  ## Build Go binary
	cp bin/config /usr/local/bin/config

#-- Docker
docker-build: ## Build Go binary
	docker build -t testconfig:latest -f dockerfile .

docker-run: docker-build ## Build Go binary
	docker run -it testconfig:latest
