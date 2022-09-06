package main

import (
	pb "github.com/gaofubao/learning-go/golang/rpc/hello/proto"
	"github.com/gaofubao/learning-go/golang/rpc/hello/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, service.NewGreeter())

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(s.Serve(lis))
}
