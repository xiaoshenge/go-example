package main

import (
	"io"
	"fmt"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"net"
	"log"
	"context"
	pb "github.com/xiaoshenge/go-example/grpc/echo/proto"
)

type server struct{}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("received: %v", in.Message)
	return &pb.EchoResponse{Message: "ack "+ in.Message}, nil
}
func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		msg := in.Message

		for index := 0; index < 3; index++ {
			if err := stream.Send(&pb.EchoResponse{Message: fmt.Sprintf("ack %s in %d", msg, index)}); err != nil {
				return err
			}
		}

	}
}

func (s *server) ServerStreamingEcho(in *pb.EchoRequest, stream pb.Echo_ServerStreamingEchoServer) error {
	msg := in.Message
	log.Printf("get msg: %s", msg)
	for index := 0; index < 3; index++ {
		if err := stream.Send(&pb.EchoResponse{Message: fmt.Sprintf("ack %s in %d", msg, index)}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.EchoResponse{Message: "bye bye"})
		}
		if err != nil {
			return err
		}
		log.Printf("get msg %s", in.Message)
	}
}

func main()  {
		lis, err := net.Listen("tcp", ":5051")
		if err != nil {
			log.Fatalf("failed to listen:%v", err)
		}

		s := grpc.NewServer()
		pb.RegisterEchoServer(s, &server{})
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve:%v", err)
		}
}