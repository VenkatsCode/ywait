package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../pb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	orderConn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	deliveryConn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))
	registerOrderService(r, orderConn)
	registerDeliveryService(r, deliveryConn)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func registerOrderService(r *gin.Engine, conn *grpc.ClientConn) {
	client := pb.NewOrderServiceClient(conn)
	r.POST("/order/place/:a", func(c *gin.Context) {
		name := c.Param("a")
		body := &pb.Order{
			OrderId:       fmt.Sprint(time.Now().Unix()),
			StoreLocation: "H3A3L4",
			Customer: &pb.CustomerInfo{
				CustomerId:       fmt.Sprintf("%s_%d", name, time.Now().Unix()),
				Name:             name,
				DeliveryLocation: "H3L1L2",
				Phone:            "+14164545447"},
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

type DevelieryInput struct {
	OrderId    string `json:"orderId"`
	DeliveryId string `json:"deliveryId"`
}

func registerDeliveryService(r *gin.Engine, conn *grpc.ClientConn) {
	client := pb.NewDeliveryServiceClient(conn)
	r.POST("/order/accept", func(c *gin.Context) {
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := string(buf[0:num])
		log.Println("input: ", reqBody)
		output := &DevelieryInput{}
		err := json.Unmarshal(buf[0:num], output)
		if err != nil {
			log.Error("error unmarshalling response ", err)
		}
		log.Println("output: ", output)
		body := &pb.DeliveryOrder{
			OrderId:    output.OrderId,
			DeliveryId: output.DeliveryId,
		}
		log.Println("Req body: ", body)
		if res, err := client.AcceptDelivery(c, body); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.POST("/order/delivered", func(c *gin.Context) {
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := string(buf[0:num])
		log.Println("input: ", reqBody)
		output := &DevelieryInput{}
		err := json.Unmarshal(buf[0:num], output)
		if err != nil {
			log.Error("error unmarshalling response ", err)
		}
		log.Println("output: ", output)
		body := &pb.DeliveryOrder{
			OrderId:    output.OrderId,
			DeliveryId: output.DeliveryId,
		}
		log.Println("Req body: ", body)
		if res, err := client.ConfirmDelivery(c, body); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}
