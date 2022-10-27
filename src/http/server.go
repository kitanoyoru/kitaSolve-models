package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "kitasolve/models/grpc"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
  solveMap map[string]pb.Solve // TODO: Change to db connection
}

// Add solve to db
func (s *Server) AddSolve(ctx context.Context, solve *pb.Solve) (*wrappers.StringValue, error) {
  s.solveMap[solve.Id] = *solve
  return &wrappers.StringValue {
    Value: fmt.Sprintf("[200] Order with Id: %v  was added successfully", solve.Id), // TODO: format http code in response 
  }, status.New(codes.OK, "").Err() 
}

// Get solve from db
func (s *Server) GetSolve(ctx context.Context, id *wrappers.StringValue) (*pb.Solve, error) {
  if solve, ok := s.solveMap[id.Value]; ok {
    return &solve, status.New(codes.OK, "").Err()
  } else {
    return nil, status.Errorf(codes.NotFound, "Solve doesn't exists: ", id)
  }
}

// Get all solves which incldue this query
func (s *Server) SearchSolves(searchQuery *wrappers.StringValue, stream pb.KitaSolveModels_SearchSolveServer) error {
  for _, value := range s.solveMap {    // TODO: Change to findAll mongodb method
    if strings.Contains(value.Content, searchQuery.Value) {
      err := stream.Send(&value)
      if err != nil {
        return fmt.Errorf("[ERROR] Error in sending stream output to client ")
      }
      break
    }
  }
  return nil
}

// Update streams
func (s *Server) UpdateSolves(stream pb.KitaSolveModels_UpdateSolveServer) error {
  for {
    solve, err := stream.Recv()
    if err == io.EOF {
      return stream.SendAndClose(&wrappers.StringValue { Value: "[200] Requested updates was processed" })
    }
    if err != nil {
      return err
    }

    s.solveMap[solve.Id] = *solve // TODO: Change to updateOne mongodb method
  }
}











