#!/bin/bash
buildNumber=`date +%Y%m%d%.%H%M%S`
go build -ldflags "-X main.buildNumber $buildNumber" -o bin/cargo_osx main.go

