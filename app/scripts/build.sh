#!/bin/bash

TRG_PKG='main'

VERSION="1.0.0"
BUILD_DATE=$(date +'%Y-%m-%d')

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags "-X $TRG_PKG.version=$VERSION -X $TRG_PKG.buildDate=$BUILD_DATE" -o ./bin/GophKeeper-windows-amd64.exe ./../app/cmd/client


# Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-X $TRG_PKG.version=$VERSION -X $TRG_PKG.buildDate=$BUILD_DATE" -o bin/GophKeeper-linux-amd64 ./../app/cmd/client


# Mac OS ARM
go build -ldflags "-X $TRG_PKG.version=$VERSION -X $TRG_PKG.buildDate=$BUILD_DATE" -o bin/GophKeeper-darwin-arm64 ./../app/cmd/client