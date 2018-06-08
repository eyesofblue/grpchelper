package main

import (
	"google.golang.org/grpc"
	"log"
    "net"
    "strconv"
	"{{.PrefixFromGoSrcPath}}/{{.DirName}}/pb"
	"{{.PrefixFromGoSrcPath}}/{{.DirName}}/svr/handler"
)

const (
    SVR_IP = "{{.SvrIp}}"
    SVR_PORT = {{.SvrPort}}
)

func main() {
	// 创建监听端口
    address := SVR_IP + ":" + strconv.FormatUint(uint64(SVR_PORT), 10)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to Listen. msg:%s", err)
	}
	// 创建grpc服务器实例
	grpcSvr := grpc.NewServer()
	// 将自己的服务器实现注册到grpc服务器上
    server := handler.NewRpcHandler()
	pb.Register{{.ProjName}}Server(grpcSvr, server)
	// 开启服务
	grpcSvr.Serve(lis)
}

