package main

import (
	"context"
	"fmt"
	// "google.golang.org/grpc/status"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"

	"../pb"
)

type orderServer struct {
}

var collection *mongo.Collection

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Product server")
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &orderServer{})
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


func (*orderServer) Get(ctx context.Context, req *pb.OrderId) (*pb.Order, error) {

	var result pb.Order
	//filter := bson.D{{"Id", req.GetId()}}
	err := collection.FindOne(ctx, req).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	return &result, nil
}

func (*orderServer) PlaceOrder(ctx context.Context, req *pb.Order) (*empty.Empty, error) {

	order := req
	order.Status = pb.Order_PLACED

	_, err := collection.InsertOne(ctx, order)

	if err != nil {
		log.Printf("Error creating order %s", err.Error())
	}

	//create delivery client and call publish order
	connDelivery, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	deliveryClient := pb.NewDeliveryServiceClient(connDelivery)
	deliveryClient.PublishOrder(ctx, req)

	return new(empty.Empty), nil
}

func (*orderServer) DeliveringOrder(ctx context.Context, req *pb.DeliveryInfo) (*empty.Empty, error) {
	filter := bson.M{"id": req.OrderId}
	_, err := collection.UpdateOne(ctx, filter,
		bson.M{"$set": bson.M{"status": pb.Order_IN_TRANSIT}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//call messaging service to send a message to the customer that the order is in transit
	//create messaging client and call send method
	connMessage, err := grpc.Dial(":7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	messageClient := pb.NewMessageServiceClient(connMessage)

	var order pb.Order
	filter1 := bson.M{"id":req.OrderId}
	err1 := collection.FindOne(ctx, filter1).Decode(&order)
	if err1 != nil {
		log.Fatalf("Cannot fetch order for ID: %v", req.OrderId)
	}

	var message pb.Message
	message.Recipient = []string{order.Customer.Phone}
	message.Message = fmt.Sprintf("Order with id %v is in transit, it is being picked up by %v and you can reach him at %v", req.OrderId, req.DeliveryPersonName, req.DeliveryPersonMobile)
	messageClient.Send(ctx, &message)
	return new(empty.Empty), nil
}


func (*orderServer) DeliveredOrder(ctx context.Context, req *pb.OrderId) (*empty.Empty, error) {

	filter := bson.M{"id": req.Id}
	_, err := collection.UpdateOne(ctx, filter,
		bson.M{"$set": bson.M{"status": pb.Order_DELIVERED}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//call messaging service to send a message to the customer that the order is delivered
	//create messaging client and call send method
	connMessage, err := grpc.Dial(":7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	messageClient := pb.NewMessageServiceClient(connMessage)

	var order pb.Order
	filter1 := bson.M{"id":req.Id}
	err1 := collection.FindOne(ctx, filter1).Decode(&order)
	if err1 != nil {
		log.Fatalf("Cannot fetch order for ID: %v", req.Id)
	}

	var message pb.Message
	message.Recipient = []string{order.Customer.Phone}
	message.Message = fmt.Sprintf("Order with id %v is in delivered", req.Id)
	messageClient.Send(ctx, &message)
	return new(empty.Empty), nil
}



