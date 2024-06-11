package service

import (
	"context"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/reservation"
	"github.com/Project_Restaurant/Restaurant-Service/storage"
	"github.com/rs/zerolog/log"
)

// ReservationService implements the reservation.ReservationServiceServer interface.
type ReservationService struct {
	stg storage.StorageI
	reservation.UnimplementedReservationServiceServer
}

// NewReservationService creates a new ReservationService.
func NewReservationService(stg storage.StorageI) *ReservationService {
	return &ReservationService{stg: stg}
}

// CreateReservation creates a new reservation.
func (s *ReservationService) CreateReservation(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error) {
	log.Info().Msg("ReservationService: CreateReservation called")

	resp, err := s.stg.Reservation().Create(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("ReservationService: Error creating reservation")
		return nil, err
	}
	return resp, nil
}

// GetReservation gets a reservation by its ID.
func (s *ReservationService) GetReservation(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationResponse, error) {
	log.Info().Msg("ReservationService: GetReservation called")

	resp, err := s.stg.Reservation().GetById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("ReservationService: Error getting reservation by ID")
		return nil, err
	}
	return resp, nil
}

// UpdateReservation updates a reservation.
func (s *ReservationService) UpdateReservation(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationResponse, error) {
	log.Info().Msg("ReservationService: UpdateReservation called")

	resp, err := s.stg.Reservation().Update(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("ReservationService: Error updating reservation")
		return nil, err
	}
	return resp, nil
}

// DeleteReservation deletes a reservation.
func (s *ReservationService) DeleteReservation(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error) {
	log.Info().Msg("ReservationService: DeleteReservation called")

	resp, err := s.stg.Reservation().Delete(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("ReservationService: Error deleting reservation")
		return nil, err
	}
	return resp, nil
}

// ListReservations lists reservations with filtering and pagination.
func (s *ReservationService) ListReservations(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error) {
	log.Info().Msg("ReservationService: ListReservations called")

	resp, err := s.stg.Reservation().ListReservations(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("ReservationService: Error listing reservations")
		return nil, err
	}
	return resp, nil
}

// CheckAvailability checks the availability of a reservation.
func (s *ReservationService) CheckAvailability(ctx context.Context, req *reservation.CheckAvailabilityRequest) (*reservation.CheckAvailabilityResponse, error) {
	log.Info().Msg("ReservationService: CheckAvailability called")

	resp, err := s.stg.Reservation().CheckAvailability(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("ReservationService: Error checking reservation availability")
		return nil, err
	}
	return resp, nil
}
