package server

import (
	"context"
	"fmt"

	pb "github.com/kitanoyoru/kitaSolve-models/pkg/grpc"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
  "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
  col *mongo.Collection 
}

func (s *Server) AddSolve(ctx context.Context, solve *pb.Solve) (*wrappers.StringValue, error) {   
  _, err := s.col.InsertOne(context.Background(), *solve)
  if err != nil {
    return &wrappers.StringValue {
      Value: fmt.Sprintf("Error while inserting new item"), // TODO: Maker handling for more specific errors
    }, status.Error(codes.Canceled, "")
  }

  return &wrappers.StringValue {
    Value: fmt.Sprintf("Order with Id: %v  was added successfully", solve.Id),
  }, status.New(codes.OK, "").Err() 
}

func (s *Server) GetSolve(ctx context.Context, id *wrappers.StringValue) (*pb.Solve, error) {
  var solve pb.Solve
  cur := s.col.FindOne(context.Background(), bson.D{{"id", id.Value}})
  
  if err := cur.Decode(&solve); err != nil {
    return nil, status.Error(codes.DataLoss, "")
  }
  
  return &solve, status.New(codes.OK, "").Err()
}

func (s *Server) DeleteSolve(ctx context.Context, id *wrappers.StringValue) (*wrappers.BoolValue, error) {
  if _, err := s.col.DeleteOne(context.Background(), bson.D{{"id", id.Value}}); err != nil {
    return &wrappers.BoolValue {
      Value: false,
    }, status.Error(codes.Internal, "")
  }

  return &wrappers.BoolValue {
    Value: true,
  }, status.New(codes.OK, "").Err()
}

