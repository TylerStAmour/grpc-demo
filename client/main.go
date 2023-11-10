package main

import (
	"context"
	"github.com/tylerstamour/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewClient(conn)
	reply, err := c.Ping(context.Background(), &proto.PingRequest{Message: "ping"})
	if err != nil {
		panic(err)
	}
	println(reply.Message)
}
