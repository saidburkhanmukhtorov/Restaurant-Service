package service

import (
	"context"

	"github.com/Diyor/Project/Restaurant/genproto/reservation"
	"github.com/Diyor/Project/Restaurant/storage"
)

// type ReservationService struct {
// 	Reservation storage.StorageI
// 	reservation.UnimplementedReservationServiceServer
// }

// func NewReservationService(reservation storage.StorageI) *ReservationService {
// 	return &ReservationService{
// 		Reservation: reservation,
// 	}
// }

// func (r *ReservationService) CreateReservation(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error) {
// 	res, err := r.Reservation.CreateReservation(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil	

// }
