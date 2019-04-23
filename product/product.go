package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"

	"../pb"
)

type productServer struct {
	gorm.Model
	savedProducts []*pb.Product
}

var collection *mongo.Collection


func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Product server")
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &productServer{})
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



func (s *productServer) Create(ctx context.Context, req *pb.Product) (*pb.Product, error) {


	product := req

	if product.Id == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Product ID cannot be empty")
	}

	if product.Description == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Product Description cannot be empty")
	}

	if product.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Product name cannot be empty")
	}

	if product.Price == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Product Price cannot be zero")
	}

	_, err := collection.InsertOne(ctx, product)

	if err != nil {
		log.Printf("Error creating Product %s", err.Error())
	}
	return product, nil
}

func (s *productServer) FindOne(ctx context.Context, req *pb.ProductId) (*pb.Product, error) {

	var result pb.Product
	//filter := bson.D{{"Id", req.GetId()}}
	collection.FindOne(ctx, req).Decode(&result)
	fmt.Printf("Found a single document: %+v\n", result)

	return &result, nil

}


func (s *productServer) FindAll(req *empty.Empty, stream pb.ProductService_FindAllServer) error {

	findOptions := options.Find()

	var productsList []*pb.Product

	cur, err := collection.Find(context.TODO(), nil, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem pb.Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		productsList = append(productsList, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", productsList)

	return status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}

func (s *productServer) Update(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *productServer) Delete(ctx context.Context, req *pb.ProductId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *productServer) Validate(ctx context.Context, req *pb.ValidateQuantity) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
