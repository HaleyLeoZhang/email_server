#!/bin/bash
gofmt -w ./middleware/jwt/*.go
gofmt -w ./*.go
gofmt -w ./caches/*.go
gofmt -w ./models/*.go
gofmt -w ./pkg/app/*.go
gofmt -w ./pkg/e/*.go
gofmt -w ./pkg/file/*.go
gofmt -w ./pkg/gredis/*.go
gofmt -w ./pkg/logging/*.go
gofmt -w ./pkg/setting/*.go
gofmt -w ./pkg/util/*.go

gofmt -w ./routers/api/comic/*.go

gofmt -w ./service/comic_service/*.go

