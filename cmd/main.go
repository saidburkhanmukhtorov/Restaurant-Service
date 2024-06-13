package main

import (
	"log"
	"net"

	"github.com/Diyor/Project/Restaurant/genproto/menu"
	"github.com/Diyor/Project/Restaurant/genproto/reservation"
	"github.com/Diyor/Project/Restaurant/genproto/restaurant"
	"github.com/Diyor/Project/Restaurant/service"
	"github.com/Diyor/Project/Restaurant/storage/postgres"
	"google.golang.org/grpc"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// s := grpc.NewServer()
	// menu.RegisterMenuServiceServer(s, service.NewMenuService(db))
	// reservation.RegisterReservationServiceServer(s, service.NewReservationService(db))
	// restaurant.RegisterRestaurantServiceServer(s, service.NewRestaurantService(db))
	// log.Printf("server listening at %v", liss.Addr())
	// if err := s.Serve(liss); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	}
}
