package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Diyor/Project/Restaurant/genproto/reservation"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Reservation struct {
	db *sql.DB
}

func NewReservation(db *sql.DB) *Reservation {
	return &Reservation{db}
}

func (r *Reservation) Create(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error) {
	id := uuid.NewString()
	query := `
	insert into reservations
	(id, user_id, restaurant_id, reservation_time, created_at)
	values($1, $2, $3, $4, $5)
	returning id, user_id, restaurant_id, reservation_time, created_at

	`
	var res reservation.Reservation
	err := r.db.QueryRowContext(ctx, query, id, req.Reservation.UserId, req.Reservation.RestaurantId, req.Reservation.ReservationTime, req.Reservation.CreatedAt).
		Scan(&res.Id, &res.UserId, &res.RestaurantId, &res.ReservationTime, &res.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &reservation.CreateReservationResponse{
		Reservation: &res,
	}, nil
}

func (r *Reservation) Update(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationResponse, error) {
	query :=
		`
	update reservations
	set user_id = $1, restaurant_id = $2, reservation_time = $3, updated_at = $4
	where id = $5
	returning id, user_id, restaurant_id, reservation_time, updated_at
	`
	res := reservation.Reservation{}
	err := r.db.QueryRowContext(ctx, query, req.Reservation.UserId, req.Reservation.RestaurantId, req.Reservation.ReservationTime, req.Reservation.UpdatedAt).
		Scan(&res.Id, &res.UserId, &res.RestaurantId, &res.ReservationTime, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &reservation.UpdateReservationResponse{
		Reservation: &res,
	}, nil
}

func (r *Reservation) Delete(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error) {
	query := `delete from reservations where id = $1`
	_, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	return &reservation.DeleteReservationResponse{}, nil
}

func (r *Reservation) Get(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationResponse, error) {
	query := `select id, user_id, restaurant_id, reservation_time, created_at, updated_at from reservations where id = $1`
	res := reservation.Reservation{}
	err := r.db.QueryRowContext(ctx, query, req.Id).
		Scan(&res.Id, &res.UserId, &res.RestaurantId, &res.ReservationTime, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &reservation.GetReservationResponse{
		Reservation: &res,
	}, nil
}

func (r *Reservation) GetAll(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			user_id,
			restaurant_id,
			reservation_time,
			status,
			created_at,
			updated_at
		FROM 
			reservations
		WHERE 1=1 
	`

	filter := ""

	if req.UserId != "" {
		filter += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.RestaurantId != "" {
		filter += fmt.Sprintf(" AND restaurant_id = $%d", count)
		args = append(args, req.RestaurantId)
		count++
	}

	if req.Status != "" {
		filter += fmt.Sprintf(" AND status = $%d", count)
		args = append(args, req.Status)
		count++
	}

	if req.StartTime != "" {
		startTime, err := time.Parse(time.RFC3339, req.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start time format: %w", err)
		}
		filter += fmt.Sprintf(" AND reservation_time >= $%d", count)
		args = append(args, startTime)
		count++
	}

	if req.EndTime != "" {
		endTime, err := time.Parse(time.RFC3339, req.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end time format: %w", err)
		}
		filter += fmt.Sprintf(" AND reservation_time <= $%d", count)
		args = append(args, endTime)
		count++
	}

	query += filter

	// Apply pagination
	if req.Limit <= 0 {
		req.Limit = 10 // Default limit
	}
	if req.Page <= 0 {
		req.Page = 1 // Default page
	}
	offset := (req.Page - 1) * req.Limit
	query += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, req.Limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error listing reservations")
		return nil, err
	}
	defer rows.Close()

	var reservations []*reservation.Reservation
	for rows.Next() {
		var (
			dbReservation   reservation.Reservation
			createdAt       time.Time
			updatedAt       time.Time
			reservationTime time.Time
		)
		err := rows.Scan(
			&dbReservation.Id,
			&dbReservation.UserId,
			&dbReservation.RestaurantId,
			&reservationTime,
			&dbReservation.Status,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning reservation row")
			return nil, err
		}
		dbReservation.CreatedAt = createdAt.Format(time.RFC3339)
		dbReservation.UpdatedAt = updatedAt.Format(time.RFC3339)
		dbReservation.ReservationTime = reservationTime.Format(time.RFC3339)
		reservations = append(reservations, &dbReservation)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over reservation rows")
		return nil, err
	}

	return &reservation.ListReservationsResponse{Reservations: reservations}, nil
}
