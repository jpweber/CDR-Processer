#!/bin/bash
buildNumber=`date +%Y%m%d%.%H%M%S`
#os x build
go build -ldflags "-X main.buildNumber $buildNumber" -o main main.go
#linux build
#env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.buildNumber $buildNumber" -o cdr_server cdr_server.go
