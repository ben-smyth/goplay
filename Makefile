.PHONY: help

.DEFAULT_GOAL := help

help: ## Show this help.
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@echo '  help    Show this help message'
	@echo '  build   Compile the project'
	@echo '  clean   Remove compiled files'
	@echo '  test    Run tests'
	@echo '  swagger Create swagger files'

build: ## Compile the project
	@echo 'Building the project...'
	# Add your build commands here

clean: ## Remove compiled files
	@echo 'Cleaning up...'
	# Add your clean commands here

test: ## Run tests
	@echo 'Running tests...'

# swagger: ## Generate swagger files
# 	@echo 'Generating swagger files...'
# 	@swagger generate spec -i ./api/oapi.json -o ./api/swagger/swagger.json --scan-models

openapi: ## Generate API Code
	@echo 'Generate API code and UI from api/oapi.yaml...'
	@oapi-codegen -package gen -generate models api/oapi.yaml > api/gen/model-gen.go
	@yq eval api/oapi.yaml -o=json > api/swagger/oapi.json
