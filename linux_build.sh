#!/bin/bash
buildNumber=`date +%Y%m%d%.%H%M%S`
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.buildNumber $buildNumber" -o bin/cargo_linux main.go
cp bin/cargo_linux dist/cargo

