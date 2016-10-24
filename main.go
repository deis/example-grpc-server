package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/deis/example-grpc-server/_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// server is used to implement WelcomeServer.
type server struct{}

// PoweredBy implements WelcomeServer
func (s *server) PoweredBy(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: fmt.Sprintf("Powered by %s\n", in.Name)}, nil
}

func main() {
	port := getenv("PORT", "5000")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %v...\n", port)
	s := grpc.NewServer()
	pb.RegisterPoweredByServer(s, &server{})
	s.Serve(lis)
}

func getenv(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}
