package main

import (
	"context"
	"github.com/tylerstamour/grpc-demo/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
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

	c := proto.NewDemoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.Ping(ctx, &proto.PingRequest{Message: "Ping!"})
	if err != nil {
		panic(err)
	}
	println(reply.Message)
}
