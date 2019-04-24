package main

import (
	"context"
	"io"
	"log"
	"testing"

	"../pb"
	"github.com/golang/protobuf/ptypes/empty"
)

var (
	sampleCart *pb.Cart
)

func TestCreateCart(t *testing.T) {
	cartsSetup()
	defer cartsShutdown()
	createdCart := createCart(t, sampleCart)
	defer deleteCart(t, sampleCart)

	if createdCart.Id != sampleCart.Id {
		t.Error("Correct cart id not created")
	}
}

func TestDeleteCart(t *testing.T) {
	cartsSetup()
	defer cartsShutdown()

	createCart(t, sampleCart)
	count := len(*getCarts(t))

	deleteCart(t, sampleCart)
	if len(*getCarts(t)) != count-1 {
		t.Error("Delete cart failed")
	}
}

func createCart(t *testing.T, cart *pb.Cart) *pb.Cart {
	createdCart, err := cartClient.Create(context.TODO(), cart)
	if err != nil {
		t.Fatal("Creating cart failed, ", err)
	}
	return createdCart
}

func deleteCart(t *testing.T, cart *pb.Cart) {
	_, err := cartClient.Delete(context.TODO(), &pb.CartId{Id: cart.Id})
	if err != nil {
		t.Fatal("Creating cart failed, ", err)
	}
}

func getCarts(t *testing.T) *[]*pb.Cart {
	stream, err := cartClient.FindAll(context.TODO(), new(empty.Empty))
	if err != nil {
		t.Fatal("Failed retrieving carts from client", err)
	}
	var carts []*pb.Cart
	for {
		cart, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error recieving cart from stream, ", err)
			continue
		}
		carts = append(carts, cart)
	}
	return &carts
}

func cartsSetup() {
	sampleCart = &pb.Cart{Id: "test_id", Description: "a cart", Status: pb.Cart_CREATED}
}

func cartsShutdown() {

}
