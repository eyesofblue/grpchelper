package main

import (
	"log"
	"github.com/eyesofblue/grpchelper/logic"
	"google.golang.org/grpc"
	"context"
    "strconv"
	"{{.GoModulePath}}/pb"
	_ "{{.GoModulePath}}/cli_tool/stub"
)

const (
    SVR_IP = "{{.SvrIp}}"
    SVR_PORT = {{.SvrPort}}
)

func main() {
	// Set up a connection to the server.
    address := SVR_IP + ":" + strconv.FormatUint(uint64(SVR_PORT), 10)
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.New{{.ProjName}}Client(conn)
    logic.ClientStub(ctx, c)
}

