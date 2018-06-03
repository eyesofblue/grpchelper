/*
 * author: jinwei
 * date  : {{.Date}}
 */

package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "{{.PathPrefix}}/{{.ModuleName}}"
)

const (
    SVR_IP = {{.SvrIp}}
    SVR_PORT = {{.SvrPort}}
)

func main() {
	// Set up a connection to the server.
    address := SVR_IP + ":" + SVR_PORT
    conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.New{{.ModuleName}}Client(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	/*
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
    */
	log.Printf("Greeting: %s", r.Message)
}
