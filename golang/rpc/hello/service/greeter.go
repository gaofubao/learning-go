package service

import (
	"context"
	"fmt"
	pb "github.com/gaofubao/learning-go/golang/rpc/hello/proto"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("hello")
	return nil, nil
}
