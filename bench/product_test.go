package main

import (
	"context"
	"io"
	"log"
	"testing"

	"../pb"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
)

var (
	sampleProduct *pb.Product
)

func TestCreateProduct(t *testing.T) {
	productsSetup()
	defer productsShutdown()

	createdProduct := createProduct(t, sampleProduct)
	defer deleteProduct(t, sampleProduct)

	if createdProduct.Id != sampleProduct.Id {
		t.Error("Correct product id not created")
	}
}

func TestDeleteProduct(t *testing.T) {
	productsSetup()
	defer productsShutdown()

	createProduct(t, sampleProduct)
	count := len(*getProducts(t))

	deleteProduct(t, sampleProduct)
	if len(*getProducts(t)) != count-1 {
		t.Error("Delete product failed")
	}
}

func createProduct(t *testing.T, product *pb.Product) *pb.Product {
	createdProduct, err := productClient.Create(context.TODO(), product)
	if err != nil {
		t.Fatal("Creating cart failed, ", err)
	}
	return createdProduct
}

func deleteProduct(t *testing.T, product *pb.Product) {
	_, err := productClient.Delete(context.TODO(), &pb.ProductId{Id: product.Id})
	if err != nil {
		t.Fatal("Creating cart failed, ", err)
	}
}

func getProducts(t *testing.T) *[]*pb.Product {
	stream, err := productClient.FindAll(context.TODO(), new(empty.Empty))
	if err != nil {
		t.Fatal("Failed retrieving carts from client", err)
	}
	var products []*pb.Product
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error recieving product from stream, ", err)
			continue
		}
		products = append(products, product)
	}
	return &products
}

func productsSetup() {
	sampleProduct = &pb.Product{
		Id:          "test_product_id",
		Name:        "Coolest product",
		Description: "a product",
		Price:       2.20,
		Quantity:    15,
		Details:     map[string]*any.Any{},
	}
}

func productsShutdown() {

}
