package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/grpc_tutorial/api/routes"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {

	flag.Parse() // initializes cli flags if used

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen to port 8080: %v", err)
	}

	s := grpc.NewServer()
	routes.SetupRoutes(s)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen to port 8080: %v", err)
	}
}
