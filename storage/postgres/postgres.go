package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Diyor/Project/Restaurant/storage"
)

type Storage struct {
	Db         *sql.DB
	Restaurant storage.RestaurantI
	Reservation storage.ReservationI
}

func ConnectDB() (*Storage, error) {
	psql := "user=postgres password=20005 dbname=restarount sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	r := NewRestaurants(db)
	
	return &Storage{
		Db:       db,
		Restaurant: r,
	}, nil
}


func (s *Storage) Restaurants() storage.RestaurantI {
	if s.Restaurant == nil {
		s.Restaurant = NewRestaurants(s.Db)
	}
	return s.Restaurant
}

