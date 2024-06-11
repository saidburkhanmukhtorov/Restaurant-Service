package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Project_Restaurant/Restaurant-Service/config"
	"github.com/Project_Restaurant/Restaurant-Service/storage"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB pgx.Conn
}

func DBConn() (*Storage, error) {
	var (
		db  *pgx.Conn
		err error
	)
	// Get postgres connection data from .env file
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	// Connecting to postgres
	db, err = pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &Storage{}, err
}

func (s *Storage) Menu() storage.MenuI {
	return nil
}
func (s *Storage) Reservation() storage.ReservationI {
	return nil
}

func (s *Storage) Restaurant() storage.RestaurantI {
	return nil
}
