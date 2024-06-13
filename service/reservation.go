package service

import (
	"context"

	"github.com/Diyor/Project/Restaurant/genproto/reservation"
	"github.com/Diyor/Project/Restaurant/storage"
)

type ReservationService struct {
	Reservation storage.StorageI
	reservation.UnimplementedReservationServiceServer
}

func NewReservationService(reservation storage.StorageI) *ReservationService {
	return &ReservationService{
		Reservation: reservation,
	}
}

func (r *ReservationService) CreateReservation(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error) {
	res, err := r.Reservation.Reservation().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (r *ReservationService) GetReservation(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationResponse, error) {
	res, err := r.Reservation.Reservation().Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) UpdateReservation(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationResponse, error) {
	res, err := r.Reservation.Reservation().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) DeleteReservation(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error) {
	res, err := r.Reservation.Reservation().Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) ListReservations(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error) {
	res, err := r.Reservation.Reservation().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) CheckAvailability(ctx context.Context, req *reservation.CheckAvailabilityRequest) (*reservation.CheckAvailabilityResponse, error) {
	res, err := r.Reservation.Reservation().CheckAvailability(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) FoodList(ctx context.Context, req *reservation.OrderFoodListReq) (*reservation.OrderFoodListRes, error) {
	res, err := r.Reservation.Reservation().FoodList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) OrderFood(ctx context.Context, req *reservation.OrderFoodReq) (*reservation.OrderFoodRes, error) {
	res, err := r.Reservation.Reservation().OrderFood(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}	
