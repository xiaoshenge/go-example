package main

import (
	"google.golang.org/grpc/resolver"
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
	pb "github.com/xiaoshenge/go-example/grpc/echo/proto"
)

func main()  {
	// passthroughConn, err := grpc.Dial(
	// 	"passthrough:///localhost:5052",
	// 	grpc.WithInsecure(),
	// )
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer passthroughConn.Close()

	// makeRPCs(passthroughConn, 10)


	exampleConn, err := grpc.Dial(
		"example:///resolver.example.grpc.io",
		grpc.WithInsecure(),
		grpc.WithBalancerName("round_robin"),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer exampleConn.Close()

	makeRPCs(exampleConn, 10)

}


func makeRPCs(conn *grpc.ClientConn, n int)  {
	client := pb.NewEchoClient(conn)
	for i:=0; i<n;i++{
		UnaryEcho(client, "this is name_resolving")
	}
}

func UnaryEcho(client pb.EchoClient, msg string)  {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message:msg})
	if err != nil {
		log.Fatalf("could not echo：%v", err)
	}

	log.Printf("echo Response: %s", r.Message)
}

// name resolver

type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build (target resolver.Target, cc resolver.ClientConn, opt resolver.BuildOption) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc: cc, 
		addrsStore: map[string][]string{
			"resolver.example.grpc.io": {"localhost:5052","localhost:5053","localhost:5054"},
		},
	}
	r.start()
	return r,nil
}

func (*exampleResolverBuilder) Scheme() string {return "example"}

type exampleResolver struct{
	target resolver.Target
	cc resolver.ClientConn
	addrsStore map[string][]string
}
func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i,s := range addrStrs{
		addrs[i] = resolver.Address{Addr:s}
	}
	r.cc.NewAddress(addrs)
}

func (*exampleResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*exampleResolver) Close()                                 {}

func init()  {
	resolver.Register(&exampleResolverBuilder{})
}