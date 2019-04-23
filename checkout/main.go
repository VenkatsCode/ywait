package main

import (
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"

	"../pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting checkout server, listening on port 8888...")
	s := grpc.NewServer()
	pb.RegisterCheckoutServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

func (s *server) Checkout(ctx context.Context, input *pb.CheckoutInput) (*empty.Empty, error) {
	cartID := input.CartId
	log.Printf("checkout cart ID is %s", cartID)

	//create cart client
	connCart, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	cartClient := pb.NewCartServiceClient(connCart)

	//create product client
	connProduct, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	productClient := pb.NewProductServiceClient(connProduct)

	// Call Cart service
	reqCart := &pb.CartId{Id: cartID}

	//get cart from cart id
	if resCart, err := cartClient.FindOne(ctx, reqCart); err == nil {
		//update cart, change status
		retrievedCart := resCart
		retrievedCart.Status = 1 //change cart status to PLACED
		cartClient.Update(ctx, retrievedCart)

		//get products and their quantities from cart
		cartProducts := retrievedCart.Products
		//for every product, update quantity
		for productID, productQuantity := range cartProducts {
			log.Printf("product id [%s] quantity [%d]", productID, productQuantity)

			// Call Product service
			reqProduct := &pb.ProductId{Id: productID}
			if resProduct, err := productClient.FindOne(ctx, reqProduct); err == nil {
				//update product, change quantity
				retrievedProduct := resProduct
				retrievedProduct.Quantity = retrievedProduct.Quantity - productQuantity
				productClient.Update(ctx, retrievedProduct)
			} else {
				log.Printf("error finding product by product id [%s]", reqProduct)
			}
		}
	} else {
		log.Printf("error finding cart by cart id [%s]", reqCart)
	}

	return new(empty.Empty), nil
}
