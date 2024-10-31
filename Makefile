sdk:
	cp ../control-plane/api/api/definitions/controlplane.yml ./definition.yml
	oapi-codegen -package=sdk \
		-generate=types,client,spec \
		-o=sdk/beamlit.go \
		-templates=./templates \
		definition.yml

build:
	goreleaser release --snapshot --clean
	cp ./dist/beamlit_darwin_arm64/beamlit ./beamlit

dev:
	alias bl="go run main.go"

.PHONY: sdk