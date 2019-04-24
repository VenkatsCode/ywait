package main

import (
	"log"
	"os"
	"testing"

	"../pb"
	"google.golang.org/grpc"
)

// grpc clients
var (
	cartClient    pb.CartServiceClient
	productClient pb.ProductServiceClient
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(":7777 Dial failed: %v", err)
	}
	cartClient = pb.NewCartServiceClient(conn)

	conn, err = grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(":9999 Dial failed: %v", err)
	}
	productClient = pb.NewProductServiceClient(conn)
}

func shutdown() {
	cartClient = nil
	productClient = nil
}
