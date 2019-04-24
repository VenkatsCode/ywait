package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../pb"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var collection *mongo.Collection
var messagingClient pb.MessageServiceClient
var orderClient pb.OrderServiceClient

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Delivery service")
	s := grpc.NewServer()
	pb.RegisterDeliveryServiceServer(s, &server{})
	reflection.Register(s)
	initDatabase()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	//create messaging client
	connMessaging, err := grpc.Dial(":7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	messagingClient = pb.NewMessageServiceClient(connMessaging)

	//create order client
	connOrder, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	orderClient = pb.NewOrderServiceClient(connOrder)
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
	collection = client.Database("go-team-go").Collection("delivery")
}

func (s *server) Register(ctx context.Context, req *pb.Delivery) (*pb.Delivery, error) {
	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error registering delivery person %s", err.Error())
	}
	return req, err
}

func (s *server) Remove(ctx context.Context, req *pb.Delivery) (*empty.Empty, error) {
	_, err := collection.DeleteOne(ctx, req)
	if err != nil {
		log.Printf("Error removing delivery person %s", err.Error())
	}
	return new(empty.Empty), err
}

func (s *server) PublishOrder(ctx context.Context, req *pb.Order) (*empty.Empty, error) {
	log.Printf("publishing order")

	//create messaging client
	connMessaging, err := grpc.Dial(":7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	messagingClient = pb.NewMessageServiceClient(connMessaging)

	findOptions := options.Find()

	cur, err := collection.Find(context.TODO(), bson.M{"status": pb.Delivery_AVAILABLE}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	availableDeliverers := []string{}

	for cur.Next(context.TODO()) {
		var elem pb.Delivery
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found element: %s\n", &elem)

		availableDeliverers = append(availableDeliverers, elem.Phone)
	}

	cur.Close(context.TODO())

	var acceptUrl string
	acceptUrl = fmt.Sprintf("localhost:3000/order?%v", req.OrderId)

	reqMessaging := &pb.Message{Message: fmt.Sprintf("New order to be picked up, click %v to accept pickup and more information", acceptUrl ), Recipients: availableDeliverers, Type: pb.Message_TEXT}
	if resMessaging, err := messagingClient.Send(ctx, reqMessaging); err == nil {
		log.Printf("response from sending message to available delivery people %v", resMessaging)
	} else {
		log.Printf("error in sending message to available delivery people %v", err)
	}

	return new(empty.Empty), err
}

func (s *server) AcceptDelivery(ctx context.Context, req *pb.DeliveryOrder) (*empty.Empty, error) {
	//update status of delivery guy to delivering
	updateFilter := bson.M{"deliveryid": req.Delivery.DeliveryId}
	_, err := collection.UpdateOne(ctx, updateFilter, bson.M{"$set": bson.M{"status": pb.Delivery_DELIVERING}})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//create order client
	connOrder, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	orderClient = pb.NewOrderServiceClient(connOrder)

	//calls DeliveringOrder to update order status
	reqOrder := &pb.DeliveryInfo{OrderId: req.OrderId, DeliveryPersonName: req.Delivery.Name, DeliveryPersonMobile: req.Delivery.Phone}
	if resOrder, err := orderClient.DeliveringOrder(ctx, reqOrder); err == nil {
		log.Printf("response from calling DeliveringOrder %v", resOrder)
	} else {
		log.Printf("error in calling DeliveringOrder %v", err)
	}

	return new(empty.Empty), err
}

func (s *server) ConfirmDelivery(ctx context.Context, req *pb.DeliveryOrder) (*empty.Empty, error) {
	//create order client
	connOrder, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	orderClient = pb.NewOrderServiceClient(connOrder)

	//update status of delivery guy to available
	updateFilter := bson.M{"deliveryid": req.Delivery.DeliveryId}
	_, err = collection.UpdateOne(ctx, updateFilter, bson.M{"$set": bson.M{"status": pb.Delivery_AVAILABLE}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//calls DeliveredOrder to update order status
	reqOrder := &pb.OrderId{Id: req.OrderId}
	if resOrder, err := orderClient.DeliveredOrder(ctx, reqOrder); err == nil {
		log.Printf("response from calling DeliveredOrder %v", resOrder)
	} else {
		log.Printf("error in calling DeliveredOrder %v", err)
	}

	return new(empty.Empty), err
}
