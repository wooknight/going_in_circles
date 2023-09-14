package main

import (
	"context"
	"log"
	"net"

	"github.com/wooknight/going_in_circles/go/tinkr/grpc_client/ratings"
	"google.golang.org/grpc"
)

var rating ratings.Rating

type server struct {
	ratings.UnimplementedRatingServer
}

func (s *server) AddRating(ctx context.Context, rating *ratings.Rating) (*ratings.RatingID, error) {
	//save rating and return UUID
	return nil, nil
}

func (s *server) GetRatings(ctx context.Context, rating *ratings.UserID) (*ratings.Ratings, error) {
	//get rating and return
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("could not listen . Error : %v\n", err)
	}
	s := grpc.NewServer()
	ratings.RegisterRatingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not start server . Error : %v\n", err)
	}
}
