package server

import (
	"fmt"
	"log"
	"context"

	pb "kitasolve/models/grpc"
	
  "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
  solveMap map[string]pb.Solve // TODO: Change to db connextion
}

func (s *Server) AddSolve(ctx context.Context, solve *pb.Solve) (*wrappers.StringValue, error) {
  s.solveMap[solve.Id] = *solve
  return &wrappers.StringValue {
    Value: fmt.Sprintf("[200] Order with Id: %v  was added successfully", solve.Id), // TODO: format http code in response 
  }, status.New(codes.OK, "").Err() 
}

func (s *Server) GetSolve(ctx context.Context, id *wrappers.StringValue) (*pb.Solve, error) {
  if solve, ok := s.solveMap[id.Value]; ok {
    return &solve, status.New(codes.OK, "").Err()
  } else {
    return nil, status.Errorf(codes.NotFound, "Solve doesn't exists: ", id)
  }
}


