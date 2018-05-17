.PHONY: build

default: build

build:
	docker run --rm \
		-w /go/src/github.com/Dataman-Cloud/healthcheck \
		-e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64  \
		-v $(shell pwd):/go/src/github.com/Dataman-Cloud/healthcheck \
		golang:1.8.1-alpine \
		sh -c "go build -v"
	docker build --tag healthcheck:latest --rm .

binary:
	go build -v 
