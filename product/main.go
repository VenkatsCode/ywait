package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"../pb"
)

type productServer struct {
}

var collection *mongo.Collection

func main() {
	lis, err := net.Listen("tcp", ":9999")
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
	collection = client.Database("go-team-go").Collection("product")
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
	err := collection.FindOne(ctx, req).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	return &result, nil

}

func (s *productServer) FindAll(req *empty.Empty, stream pb.ProductService_FindAllServer) error {
	findOptions := options.Find()

	var productsList []*pb.Product

	cur, err := collection.Find(context.TODO(), bson.M{}, findOptions)
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

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", productsList)

	cur.Close(context.TODO())

	for _, feature := range productsList {
		if err := stream.Send(feature); err != nil {
			return err
		}
	}

	return nil
}

func (s *productServer) Update(ctx context.Context, req *pb.Product) (*pb.Product, error) {

	filter := bson.M{"id": req.Id}
	updateResult, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"id": req.Id, "name": req.Name, "description": req.Description, "details": req.Details, "price": req.Price, "quantity": req.Quantity}})

	if err != nil {
		log.Fatal("Error updating Product %s", err.Error())
	} else {
		log.Println(updateResult)
	}

	fmt.Printf("Updated: %+v\n", updateResult)
	return req, nil

}

func (s *productServer) Delete(ctx context.Context, req *pb.ProductId) (*empty.Empty, error) {

	deleteResult, err := collection.DeleteOne(context.TODO(), req)
	if err != nil {
		log.Fatal("Error deleting Product %s", err.Error())
	}
	fmt.Printf("Deleted %v product\n", deleteResult.DeletedCount)

	return new(empty.Empty), nil
}

func (s *productServer) Validate(ctx context.Context, req *pb.ValidateQuantity) (*empty.Empty, error) {

	var result pb.Product
	filter := bson.D{{"Id", req.GetId()}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal("No product found wiht Id: %+v", result.GetQuantity())
	}
	if result.GetQuantity() < req.GetQuantity() {
		return nil, grpc.Errorf(codes.InvalidArgument, "Product quantity cannot be more than: %+v", result.GetQuantity())
	}
	return nil, status.Errorf(codes.OK, "Validated")
}
