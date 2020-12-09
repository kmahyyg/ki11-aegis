#!/bin/bash
go mod download
go generate cmd/main.go
CGO_ENABLED=0 go build -trimpath -ldflags '-s -w' -o 'killaegis' cmd/main.go
upx -9 -v -o killaegis.upx killaegis