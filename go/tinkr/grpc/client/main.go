package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wooknight/going_in_circles/go/tinkr/grpc/product_info"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080")
	if err != nil {
		log.Fatalf("could not connect to server on 8080 . Error : %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: "Keri", Description: "Gorgeous"})
	if err != nil {
		log.Fatalf("Could not create product .Err : %v", err)
	}
	m, err := c.GetProduct(ctx, &pb.ProductID{Value: r})
	if err != nil {
		log.Fatalf("Could not get product .Err : %v", err)
	}
	log.Printf("Successfully got everything . %+v", m)

}
