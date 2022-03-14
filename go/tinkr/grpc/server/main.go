package main

import (
	"context"
	"log"
	"net"

	//https://linguinecode.com/post/how-to-import-local-files-packages-in-golang
	pb "github.com/wooknight/going_in_circles/go/tinkr/grpc/product_info"
	"google.golang.org/grpc"
)

var prod pb.Product

type server struct {
	pb.UnimplementedProductInfoServer
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	log.Println("My grpc server is working . Creating a new product")
	prod = *in
	return &pb.ProductID{Value: "1"}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	log.Println("My grpc server is working. Sending a product ID")
	return &prod, nil
}
func main() {
	lis, _ := net.Listen("tcp", ":8080")
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on port :8080 - Error : %v ", err)
	}
}
