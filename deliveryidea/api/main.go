package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"../pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	order, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	// delivery, err := grpc.Dial(":8080", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Dial failed: %v", err)
	// }

	r := gin.Default()
	registerOrderService(r, order)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func registerOrderService(r *gin.Engine, conn *grpc.ClientConn) {
	client := pb.NewOrderServiceClient(conn)
	r.POST("/order/:a", func(c *gin.Context) {
		name := c.Param("a")
		body := &pb.Order{
			OrderId:       fmt.Sprint(time.Now().Unix()),
			StoreLocation: "H3A3L4",
			Customer: &pb.CustomerInfo{
				CustomerId:       fmt.Sprintf("%s_%d", name, time.Now().Unix()),
				Name:             name,
				DeliveryLocation: "H3L1L2",
				Phone:            "+15145156646"},
		}
		if res, err := client.PlaceOrder(c, body); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}
