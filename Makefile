.PHONY: build clean tool lint help

all: build

build:
	@echo "App is creating. Please wait ..."
	@go build -o email_server -v . 
	@echo "App is created"
	@echo "Copy conf/app.ini To /usr/local/etc/mail.ops.hlzblog.top.ini ---DOING"
	@cp conf/app.ini /usr/local/etc/mail.ops.hlzblog.top.ini
	@echo "Copy ---DONE"

ini:
	@cp conf/app.ini /usr/local/etc/mail.ops.hlzblog.top.ini

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf email_server
	go clean -i .

test:
	@echo "Test --- START"
	@go test -v service/email_service/*.go
	@go test -v pkg/queue/*.go
	@go test -v pkg/util/*.go
	@echo "Test --- END"


help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
