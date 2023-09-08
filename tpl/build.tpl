#!/bin/bash

#protoc -I . --go_out=plugins=grpc:. pb/*.proto
# -I是设置搜索路径 --go-grpc_opt=paths=source_relative表明生成文件输出使用相对路径,按照go_out的基目录+option go_package的路径生成
protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto

if [ ! -d "bin" ]; then
    mkdir ./bin
fi

cd ./bin

go build -o svrmain ../svr/svr_main.go
go build -o clitool ../cli_tool/cli_tool_main.go
