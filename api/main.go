package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../pb"
	"../product_http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	r := gin.Default()
	registerGCD(r, conn)
	registerCartService(r, conn)
	registerProductServiceServer(r, conn)
	registerProductServiceHttpServer(r)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func registerCartService(r *gin.Engine, conn *grpc.ClientConn) {
	cartClient := pb.NewCartServiceClient(conn)
	r.GET("/cart/:a", func(c *gin.Context) {
		cartID := c.Param("a")
		// &pb.Cart{Description: "best cart", Products: map[string]int32{"code_1": 1, "code_2": 2}}
		req := &pb.CartId{Id: cartID}
		if res, err := cartClient.FindOne(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.POST("/cart/:a", func(c *gin.Context) {
		desc := c.Param("a")
		// &pb.Cart{Description: "best cart", Products: map[string]int32{"code_1": 1, "code_2": 2}}
		req := &pb.Cart{Description: desc}
		if res, err := cartClient.Create(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}

func registerProductServiceServer(r *gin.Engine, conn *grpc.ClientConn) {
	productClient := pb.NewProductServiceClient(conn)
	r.GET("/product/:a", func(c *gin.Context) {
		productID := c.Param("a")

		req := &pb.ProductId{Id: productID}
		if res, err := productClient.FindOne(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.POST("/product/:a", func(c *gin.Context) {
		desc := c.Param("a")
		req := &pb.Product{Description: desc}
		if res, err := productClient.Create(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}

func registerProductServiceHttpServer(r *gin.Engine) {

	r.GET("/producthttp/:a", func(c *gin.Context) {
		productID := c.Param("a")
		if res, err := product_http.FindOneHttp(c, productID); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(*res),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}

func registerGCD(r *gin.Engine, conn *grpc.ClientConn) {
	gcdClient := pb.NewGCDServiceClient(conn)

	r.GET("/gcd/:a/:b", func(c *gin.Context) {
		// Parse parameters
		a, err := strconv.ParseUint(c.Param("a"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		b, err := strconv.ParseUint(c.Param("b"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
			return
		}
		// Call GCD service
		req := &pb.GCDRequest{A: a, B: b}
		if res, err := gcdClient.Compute(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}