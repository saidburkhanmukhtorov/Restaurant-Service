package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Diyor/Project/Restaurant/genproto/restaurant"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Restaurants struct {
	db *sql.DB
}

func NewRestaurants(db *sql.DB) *Restaurants {
	return &Restaurants{db: db}
}

func (r *Restaurants) Create(ctx context.Context, req *restaurant.CreateRestaurantRequest) (*restaurant.CreateRestaurantResponse, error) {
	id := uuid.New().String()
	query :=
		`
	INSERT INTO restaurants (id, name, address, phone_number, description)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, name, address, phone_number, description, created_at, updated_at
	`
	res := restaurant.Restaurant{}
	var createdAt, updatedAt time.Time
	err := r.db.QueryRowContext(ctx, query, id, req.Restaurant.Name, req.Restaurant.Address, req.Restaurant.PhoneNumber, req.Restaurant.Description).
		Scan(&res.Id, &res.Name, &res.Address, &res.PhoneNumber, &res.Description, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	res.CreatedAt = createdAt.Format(time.RFC3339)
	res.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &restaurant.CreateRestaurantResponse{Restaurant: &res}, nil
}

func (r *Restaurants) Get(ctx context.Context, req *restaurant.GetRestaurantRequest) (*restaurant.GetRestaurantResponse, error) {
	res := restaurant.Restaurant{}
	var createdAt, updatedAt time.Time

	query := `
	SELECT id, name, address, phone_number, description, created_at, updated_at
	FROM restaurants
	WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.QueryRowContext(ctx, query, req.Id).
		Scan(&res.Id, &res.Name, &res.Address, &res.PhoneNumber, &res.Description, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	res.CreatedAt = createdAt.Format(time.RFC3339)
	res.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &restaurant.GetRestaurantResponse{Restaurant: &res}, nil
}

func (r *Restaurants) Update(ctx context.Context, req *restaurant.UpdateRestaurantRequest) (*restaurant.UpdateRestaurantResponse, error) {
	query := `
	UPDATE restaurants
	SET name = $1, address = $2, phone_number = $3, description = $4, updated_at = NOW()
	WHERE id = $5
	RETURNING id, name, address, phone_number, description, created_at, updated_at
	`
	res := restaurant.Restaurant{}
	var createdAt, updatedAt time.Time
	err := r.db.QueryRowContext(ctx, query, req.Restaurant.Name, req.Restaurant.Address, req.Restaurant.PhoneNumber, req.Restaurant.Description, req.Restaurant.Id).
		Scan(&res.Id, &res.Name, &res.Address, &res.PhoneNumber, &res.Description, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	res.CreatedAt = createdAt.Format(time.RFC3339)
	res.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &restaurant.UpdateRestaurantResponse{Restaurant: &res}, nil
}

func (r *Restaurants) Delete(ctx context.Context, req *restaurant.DeleteRestaurantRequest) (*restaurant.DeleteRestaurantResponse, error) {
	query := `
	UPDATE restaurants
	SET deleted_at = NOW()
	WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete restaurant: %w", err)
	}
	return &restaurant.DeleteRestaurantResponse{}, nil
}

func (r *Restaurants) GetAll(ctx context.Context, req *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT id, name, address, phone_number, description, created_at, updated_at
		FROM restaurants
		WHERE deleted_at IS NULL
	`
	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count)
		args = append(args, "%"+req.Name+"%")
		count++
	}

	if req.Address != "" {
		filter += fmt.Sprintf(" AND address ILIKE $%d", count)
		args = append(args, "%"+req.Address+"%")
		count++
	}

	query += filter

	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.Limit
	query += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, req.Limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error listing restaurants")
		return nil, err
	}
	defer rows.Close()

	var restaurants []*restaurant.Restaurant
	for rows.Next() {
		var dbRestaurant restaurant.Restaurant
		var createdAt, updatedAt time.Time
		err := rows.Scan(&dbRestaurant.Id, &dbRestaurant.Name, &dbRestaurant.Address, &dbRestaurant.PhoneNumber, &dbRestaurant.Description, &createdAt, &updatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning restaurant row")
			return nil, err
		}
		dbRestaurant.CreatedAt = createdAt.Format(time.RFC3339)
		dbRestaurant.UpdatedAt = updatedAt.Format(time.RFC3339)
		restaurants = append(restaurants, &dbRestaurant)
	}

	return &restaurant.ListRestaurantsResponse{Restaurants: restaurants}, nil
}
