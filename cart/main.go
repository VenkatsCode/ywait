package main

import (
	"fmt"
	"log"
	"net"

	"../pb"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
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
	collection = client.Database("go-team-go").Collection("cart")
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
	//return nil

	findOptions := options.Find()

	var cartsList []*pb.Cart

	cur, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem pb.Cart
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		cartsList = append(cartsList, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", cartsList)

	cur.Close(context.TODO())

	for _, feature := range cartsList {
		if err := in.Send(feature); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) Update(ctx context.Context, req *pb.Cart) (*pb.Cart, error) {
	filter := bson.M{"id": req.Id}
	_, err := collection.UpdateOne(ctx, filter,
		bson.M{"$set": bson.M{"id": req.Id, "description": req.Description, "products": req.Products, "total_cost": req.TotalCost, "total_items": req.TotalItems, "status": req.Status}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return req, nil
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
	} else {
		log.Printf("found cart with id %s", cart.Id)
	}

	log.Printf("product id: %s", in.ProductId)
	cart.Products[in.ProductId] += in.Quantity
	/*err = validate(ctx, in.ProductId, cart.Products[in.ProductId])
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		log.Printf("validated successfully")
	}*/

	updateFilter := bson.M{"id": in.CartId}
	_, err = collection.UpdateOne(ctx, updateFilter, bson.M{"$set": bson.M{"products": cart.Products}})

	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		log.Printf("updated successfully")
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

	if cart.Products[in.ProductId] <= in.Quantity {
		delete(cart.Products, in.ProductId)
	} else {
		cart.Products[in.ProductId] -= in.Quantity
	}

	updateFilter := bson.M{"id": in.CartId}
	_, err = collection.UpdateOne(ctx, updateFilter, bson.M{"$set": bson.M{"products": cart.Products}})

	//_, err = collection.UpdateOne(ctx, filter, cart)
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
