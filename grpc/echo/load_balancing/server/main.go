package main

import (
	"log"
	"context"
	"net"
	"google.golang.org/grpc"
	"fmt"
	"flag"
	pb "github.com/xiaoshenge/go-example/grpc/echo/proto"
)
var (
	address = flag.String("address", ":5050", "server address")
)

type server struct{}
func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("received: %v", in.Message)
	return &pb.EchoResponse{Message: "ack "+ in.Message}, nil
}
func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	return nil
}
func (s *server) ServerStreamingEcho(in *pb.EchoRequest, stream pb.Echo_ServerStreamingEchoServer) error {
	return nil
}
func (s *server) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	return nil
}

func main()  {
	flag.Parse()
	lis, err := net.Listen("tcp", *address)
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	s.Serve(lis)
}