package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/metadata"

	pb "github.com/xiaoshenge/go-example/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	md := metadata.New(map[string]string{"a": "a"})
	callCtx := metadata.NewOutgoingContext(ctx, md)
	// grpc 通过ctx传值只能通过metadata传，通过手动ctx.withValue 的值无法传递
	callCtx = context.WithValue(callCtx, "tt", "ttt")
	resp, err := client.SayHello(callCtx, &pb.HelloRequest{Name: "Dr. Seuss"})
	if err != nil {
		fmt.Println("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md)
	}

	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
