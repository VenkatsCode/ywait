package main

import (
	"log"
	"net"

	"../pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

type CartModel struct {
	gorm.Model
	// pb.Cart
	Cart *pb.Cart
}

var db *gorm.DB

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
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database \n" + err.Error())
	}
	db.AutoMigrate(&CartModel{})
}

func (s *server) Create(ctx context.Context, in *pb.Cart) (*pb.Cart, error) {
	cartModel := &CartModel{Cart: in}
	db.Create(cartModel)
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
