#!/bin/bash

protoc -I . --go_out=plugins=grpc:. pb/*.proto

if [ ! -d "bin" ]; then
    mkdir ./bin
fi

cd ./bin

go build -o svrmain ../svr/svr_main.go
go build -o clitool ../cli_tool/cli_tool_main.go
