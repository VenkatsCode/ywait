package main

import (
	"fmt"
	"log"
	"net"

	"../pb"
	"github.com/golang/protobuf/ptypes/empty"
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
	_, err := collection.InsertOne(ctx, in)
	if err != nil {
		log.Printf("Error creating cart %s", err.Error())
	}
	return in, nil
}

func (s *server) FindOne(ctx context.Context, in *pb.CartId) (*pb.Cart, error) {
	result := collection.FindOne(ctx, in)
	var x pb.Cart
	err := result.Decode(&x)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &x, nil
}

func (s *server) FindAll(_ *empty.Empty, in pb.CartService_FindAllServer) error {
	// cursor, err := collection.Find(nil, nil)
	// var x pb.Cart
	// err := result.Decode(&x)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return nil
}

func (s *server) Update(ctx context.Context, in *pb.Cart) (*pb.Cart, error) {
	var filter pb.Cart
	filter.Id = in.Id
	_, err := collection.UpdateOne(ctx, filter, in)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return in, nil
}

func (s *server) Delete(ctx context.Context, in *pb.CartId) (*empty.Empty, error) {
	_, err := collection.DeleteOne(ctx, in)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil
}

func (s *server) AddToCart(ctx context.Context, in *pb.CartQuantity) (*pb.Cart, error) {
	filter := pb.CartId{Id: in.CartId}
	cart, err := s.FindOne(ctx, &filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	cart.Products[in.ProductId] += in.Quantity
	err = validate(ctx, in.ProductId, cart.Products[in.ProductId])
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = collection.UpdateOne(ctx, filter, cart)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cart, nil
}

func (s *server) RemoveFromCart(ctx context.Context, in *pb.CartQuantity) (*pb.Cart, error) {
	filter := pb.CartId{Id: in.CartId}
	cart, err := s.FindOne(ctx, &filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if cart.Products[in.ProductId] >= in.Quantity {
		delete(cart.Products, in.ProductId)
	} else {
		cart.Products[in.ProductId] -= in.Quantity
	}

	_, err = collection.UpdateOne(ctx, filter, cart)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cart, nil
}

func validate(ctx context.Context, productID string, quantity int32) error {
	connProduct, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	productClient := pb.NewProductServiceClient(connProduct)
	_, err = productClient.Validate(ctx, &pb.ValidateQuantity{Id: productID, Quantity: quantity})
	return err
}
