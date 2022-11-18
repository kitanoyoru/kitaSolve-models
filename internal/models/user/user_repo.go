package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
  usersCollection = "users"
)

type UserRepo struct {
  col *mongo.Collection
}

func NewUserRepo(db mongo.Database) *UserRepo {
  return &UserRepo {
    col: db.Collection(usersCollection)
  }
}

func (r *UserRepo) CreateUser(ctx context.Context, user User) error {
  _, err := r.col.InsertOne(ctx, user) // TODO: Make error handling more specific
  return nil
}

func (r *UserRepo) GetByCreds(ctx context.Context, email, password string) (User, error) {
  var user User
  if err := r.col.FindOne(ctx, bson.M {"email": email, "password": password}).Decode(&user); err != nil {
    return User{}, err
  }

  return user, nil
}

