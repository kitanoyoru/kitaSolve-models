package main

import (
	"fmt"
	"log"
	"net"

	pb "kitasolve/models/grpc"
	sv "kitasolve/models/http"
	
  "google.golang.org/grpc"
)

const (
  port = 5000
)

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("Failed to listen %v port: %v", port, err)
  }

  s := grpc.NewServer()
  pb.RegisterKitaSolveModelsServer(s, &sv.Server{})

  if err := s.Serve(list); err != nil {
    log.Fatalf("Failed to serve on %v port : %v", port, err)
  }
}
