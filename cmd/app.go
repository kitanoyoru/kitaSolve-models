package main

import (
	"log"
	"net"

	pb "kitasolve/models/grpc"
	sv "kitasolve/models/http"
  "kitasolve/models/db"
	
  "google.golang.org/grpc"
)

const (
  port = 5000
)

func main() {
  // create tcp socket
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("Failed to listen %v port: %v", port, err)
  }
  
  // get needed mongo collection
  conn := db.ConnectMongo()
  col := conn.Database("kitasolve").Collection("solve")
  defer db.DisconnectMongo(conn)
  
  // grpc server
  s := grpc.NewServer()
  pb.RegisterKitaSolveModelsServer(s, &sv.Server{
    col,
  })
  
  // serve
  if err := s.Serve(lis); err != nil {
    log.Fatalf("Failed to serve on %v port : %v", port, err)
  }
}
