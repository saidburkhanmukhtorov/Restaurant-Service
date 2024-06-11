package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/reservation"
	"github.com/Project_Restaurant/Restaurant-Service/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func newTestReservation(t *testing.T) *postgres.ReservationDb {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"reservation",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.ReservationDb{Db: db}
}

func createTestReservation() *reservation.CreateReservationRequest {
	reservationTime := time.Now().Format(time.RFC3339) // Example: Reservation for tomorrow
	return &reservation.CreateReservationRequest{
		Reservation: &reservation.Reservation{
			UserId:          uuid.NewString(),
			RestaurantId:    "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			ReservationTime: reservationTime,
			Status:          "PENDING",
		},
	}
}

func TestCreateReservation(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	assert.NotEmpty(t, createdReservation.Reservation.Id)
	assert.Equal(t, testReservation.Reservation.UserId, createdReservation.Reservation.UserId)
	assert.Equal(t, testReservation.Reservation.RestaurantId, createdReservation.Reservation.RestaurantId)
	assert.Equal(t, testReservation.Reservation.ReservationTime, createdReservation.Reservation.ReservationTime)
	assert.Equal(t, testReservation.Reservation.Status, createdReservation.Reservation.Status)
}

func TestGetReservationById(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	getReservation, err := rDb.GetById(context.Background(), &reservation.GetReservationRequest{Id: createdReservation.Reservation.Id})
	if err != nil {
		t.Fatalf("Error getting reservation by ID: %v", err)
	}

	assert.Equal(t, createdReservation.Reservation.Id, getReservation.Reservation.Id)
	assert.Equal(t, createdReservation.Reservation.UserId, getReservation.Reservation.UserId)
	assert.Equal(t, createdReservation.Reservation.RestaurantId, getReservation.Reservation.RestaurantId)
	assert.Equal(t, createdReservation.Reservation.ReservationTime, getReservation.Reservation.ReservationTime)
	assert.Equal(t, createdReservation.Reservation.Status, getReservation.Reservation.Status)
}

func TestUpdateReservation(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	updatedReservationReq := &reservation.UpdateReservationRequest{
		Reservation: &reservation.Reservation{
			Id:              createdReservation.Reservation.Id,
			UserId:          uuid.NewString(), // Change the user ID
			RestaurantId:    createdReservation.Reservation.RestaurantId,
			ReservationTime: createdReservation.Reservation.ReservationTime,
			Status:          "CONFIRMED", // Change the status
		},
	}

	updatedReservation, err := rDb.Update(context.Background(), updatedReservationReq)
	if err != nil {
		t.Fatalf("Error updating reservation: %v", err)
	}

	assert.Equal(t, updatedReservationReq.Reservation.Id, updatedReservation.Reservation.Id)
	assert.Equal(t, updatedReservationReq.Reservation.UserId, updatedReservation.Reservation.UserId)
	assert.Equal(t, updatedReservationReq.Reservation.Status, updatedReservation.Reservation.Status)
	// ... other assertions
}

func TestDeleteReservation(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	_, err = rDb.Delete(context.Background(), &reservation.DeleteReservationRequest{Id: createdReservation.Reservation.Id})
	if err != nil {
		t.Fatalf("Error deleting reservation: %v", err)
	}

	_, err = rDb.GetById(context.Background(), &reservation.GetReservationRequest{Id: createdReservation.Reservation.Id})
	assert.ErrorIs(t, err, postgres.ErrReservationNotFound)
}

func TestListReservations(t *testing.T) {
	rDb := newTestReservation(t)

	// Create test reservations
	restaurantID1 := "345e1148-678e-4ec7-9fab-d55df1e9cb54"
	userID1 := uuid.NewString()
	reservationTime1 := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	reservationTime2 := time.Now().Add(48 * time.Hour).Format(time.RFC3339)

	testReservations := []*reservation.CreateReservationRequest{
		{Reservation: &reservation.Reservation{UserId: userID1, RestaurantId: restaurantID1, ReservationTime: reservationTime1, Status: "PENDING"}},
		{Reservation: &reservation.Reservation{UserId: userID1, RestaurantId: restaurantID1, ReservationTime: reservationTime2, Status: "CONFIRMED"}},
		{Reservation: &reservation.Reservation{UserId: uuid.NewString(), RestaurantId: "60e83378-4e69-4e96-b24d-87c753363725", ReservationTime: reservationTime1, Status: "PENDING"}},
	}

	for _, tr := range testReservations {
		_, err := rDb.Create(context.Background(), tr)
		if err != nil {
			t.Fatalf("Error creating test reservation: %v", err)
		}
	}

	t.Run("ListReservations without filters", func(t *testing.T) {
		resp, err := rDb.ListReservations(context.Background(), &reservation.ListReservationsRequest{})
		if err != nil {
			t.Fatalf("Error listing reservations: %v", err)
		}
		assert.GreaterOrEqual(t, len(resp.Reservations), len(testReservations))
	})

	t.Run("Filter by user ID", func(t *testing.T) {
		resp, err := rDb.ListReservations(context.Background(), &reservation.ListReservationsRequest{UserId: userID1})
		if err != nil {
			t.Fatalf("Error listing reservations by user ID: %v", err)
		}
		assert.Equal(t, 2, len(resp.Reservations))
	})

	t.Run("Filter by restaurant ID", func(t *testing.T) {
		resp, err := rDb.ListReservations(context.Background(), &reservation.ListReservationsRequest{RestaurantId: restaurantID1})
		if err != nil {
			t.Fatalf("Error listing reservations by restaurant ID: %v", err)
		}
		assert.LessOrEqual(t, 2, len(resp.Reservations))
	})
}

func TestCheckAvailability_AvailableBeforeExistingReservation(t *testing.T) {
	rDb := newTestReservation(t)

	// Create a test reservation
	restaurantID := "e547d0d3-a804-4dba-921a-106b666c304b"
	userID := uuid.NewString()
	reservationTime := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	_, err := rDb.Create(context.Background(), &reservation.CreateReservationRequest{
		Reservation: &reservation.Reservation{
			UserId:          userID,
			RestaurantId:    restaurantID,
			ReservationTime: reservationTime,
			Status:          "PENDING",
		},
	})
	if err != nil {
		t.Fatalf("Error creating test reservation: %v", err)
	}

	// Test an available time slot one hour before the existing reservation
	reqTime := time.Now().Add(22 * time.Hour).Format(time.RFC3339)
	resp, err := rDb.CheckAvailability(context.Background(), &reservation.CheckAvailabilityRequest{
		RestaurantId:    restaurantID,
		ReservationTime: reqTime,
	})

	assert.NoError(t, err)
	log.Println(resp.Available)
	assert.True(t, resp.Available, "Time slot should be available")
}
