package product_http

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	//"net"

	//"github.com/golang/protobuf/ptypes/empty"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/reflection"
	//"google.golang.org/grpc/status"
	//
	//"google.golang.org/grpc"
	"../pb"
)

var collection *mongo.Collection

func init() {
	initDatabase()
}


func initDatabase() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:nimda@cluster0-vmhnr.mongodb.net/test?retryWrites=true")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Database("go-team-go").Collection("product")
}

func FindOneHttp(ctx context.Context, req string) (*pb.Product, error) {
log.Println("test: "+req)
	var result pb.Product
	filter := bson.M{"id": req}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	return &result, nil

}