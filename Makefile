sdk-dev:
	cp ../control-plane/api/api/definitions/controlplane.yml ./definition.yml
	oapi-codegen -package=beamlit \
		-generate=types,client,spec \
		-o=cmd/beamlit/beamlit.go \
		-templates=./templates \
		definition.yml

sdk:
	oapi-codegen -package=beamlit \
		-generate=types,client,spec \
		-o=cmd/beamlit/beamlit.go \
		-templates=./templates \
		definition.yml

run:
	go run main.go

build:
	go build -o ./beamlit main.go
	tar -czvf beamlit-cli.tar.gz ./beamlit

install:
	go build -o $(shell go env GOPATH)/bin/beamlit main.go
	ln -s $(shell go exnv GOPATH)/bin/beamlit $(shell go env GOPATH)/bin/bl

dev:
	alias bl="go run main.go"