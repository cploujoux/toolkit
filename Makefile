sdk-dev:
	cp ../control-plane/api/api/definitions/controlplane.yml ./definition.yml
	oapi-codegen -package=sdk \
		-generate=types,client,spec \
		-o=sdk/beamlit.go \
		-templates=./templates \
		definition.yml

sdk:
	oapi-codegen -package=sdk \
		-generate=types,client,spec \
		-o=sdk/beamlit.go \
		-templates=./templates \
		definition.yml

run:
	go run main.go

build:
	go build -o ./dist/beamlit cli/main.go

install:
	go build -o $(shell go env GOPATH)/bin/beamlit cli/main.go
	ln -s $(shell go exnv GOPATH)/bin/beamlit $(shell go env GOPATH)/bin/bl

dev:
	alias bl="go run cli/main.go"