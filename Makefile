sdk:
	cp ../controlplane/api/api/definitions/controlplane.yml ./definition.yml
	oapi-codegen -package=sdk \
		-generate=types,client,spec \
		-o=sdk/beamlit.go \
		-templates=./templates/go \
		definition.yml

sdk-python:
	openapi-python-client generate \
		--path=definition.yml \
		--output-path=./tmp-sdk-python \
		--overwrite \
		--custom-template-path=./templates/python \
		--config=./config/openapi-python-client.yml
	cp -r ./tmp-sdk-python/beamlit/* ./sdk-python/src/beamlit/
	rm -rf ./tmp-sdk-python

sdk-ts:
	npx @hey-api/openapi-ts -i ./definition.yml -o sdk-ts -c @hey-api/client-fetch

build:
	goreleaser release --snapshot --clean
	cp ./dist/beamlit_darwin_arm64/beamlit ./beamlit

dev:
	alias bl="go run main.go"

doc:
	rm -rf docs
	go run main.go docs --format=markdown --output=docs
	rm docs/beamlit_completion_zsh.md docs/beamlit_completion_bash.md

lint:
	golangci-lint run

install:
	uv pip install openapi-python-client

.PHONY: sdk sdk-python sdk-ts