package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client{
  if err := godotenv.Load(); err != nil {
    log.Println("[ERROR] Env file doesn't exists")
  }

  uri := os.Getenv("MONGODB_URI")
  if uri == "" {
    log.Fatal("[ERROR] MongoDB uri doesn't correct")
  }

  client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
  if err != nil {
    panic(err)
  }
  
  return client 
}

func DisconnectMongo(client *mongo.Client) {
  if err := client.Disconnect(context.TODO()); err != nil {
    panic(err)
  }
}


