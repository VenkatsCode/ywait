package main

import (
	"fmt"
	"log"
	"net"

	"../pb"
	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var collection *mongo.Collection

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Cart service")
	s := grpc.NewServer()
	pb.RegisterCartServiceServer(s, &server{})
	reflection.Register(s)
	initDatabase()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
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
	collection = client.Database("test").Collection("test")
}

func (s *server) Create(ctx context.Context, in *pb.Cart) (*pb.Cart, error) {
	cartModel := in
	_, err := collection.InsertOne(ctx, cartModel)
	if err != nil {
		log.Printf("Error creating cart %s", err.Error())
	}
	return in, nil
}

func (s *server) FindOne(ctx context.Context, in *pb.CartId) (*pb.Cart, error) {
	return &pb.Cart{Id: "code_1", Description: "Hi"}, nil
}

func (s *server) FindAll(_ *empty.Empty, in pb.CartService_FindAllServer) error {
	return nil
}

func (s *server) Update(ctx context.Context, in *pb.Cart) (*pb.Cart, error) {
	return nil, nil
}

func (s *server) Delete(ctx context.Context, in *pb.CartId) (*empty.Empty, error) {
	return nil, nil
}

func (s *server) AddToCart(ctx context.Context, in *pb.ValidateQuantity) (*pb.Cart, error) {
	return nil, nil
}

func (s *server) RemoveFromCart(ctx context.Context, in *pb.ValidateQuantity) (*pb.Cart, error) {
	return nil, nil
}
