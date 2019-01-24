package main

import (
	"io"
	"time"
	"context"
	"log"
	"google.golang.org/grpc"
	pb "github.com/xiaoshenge/go-example/grpc/echo/proto"
)

func main()  {

	conn, err := grpc.Dial(":5051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	// UnaryEcho(client)

	// BidirectionalStreamingEcho(client)
	ServerStreamingEcho(client)
	// ClientStreamingEcho(client)

}

func UnaryEcho(client pb.EchoClient)  {
	

	msg := "hi echo server"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message:msg})
	if err != nil {
		log.Fatalf("could not echo：%v", err)
	}

	log.Printf("echo Response: %s", r.Message)
}

func BidirectionalStreamingEcho(client pb.EchoClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	stream, err := client.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}

	waitc := make(chan struct{})

	go func ()  {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("get response: %s", in.Message)
		}
	}()

	stream.Send(&pb.EchoRequest{Message: "hi"})
	stream.Send(&pb.EchoRequest{Message: "hi hi..."})
	stream.CloseSend()

	<-waitc
}

func ServerStreamingEcho(client pb.EchoClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	stream, err := client.ServerStreamingEcho(ctx, &pb.EchoRequest{Message: "hi"})
	if err != nil {
		log.Fatalf("%v", err)
	}

	waitc := make(chan struct{})

	go func ()  {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("get response: %s", in.Message)
		}
	}()
	stream.CloseSend()

	<-waitc
}

func ClientStreamingEcho(client pb.EchoClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	stream, err := client.ClientStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}

	stream.Send(&pb.EchoRequest{Message: "hi"})
	stream.Send(&pb.EchoRequest{Message: "hi hi ..."})

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("receive: %v", reply.Message)
}