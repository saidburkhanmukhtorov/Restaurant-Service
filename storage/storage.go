package storage

import (
	"context"

	"github.com/Diyor/Project/Restaurant/genproto/menu"
	"github.com/Diyor/Project/Restaurant/genproto/reservation"
	"github.com/Diyor/Project/Restaurant/genproto/restaurant"
)

type StorageI interface {
	Menu() MenuI
	Reservation() ReservationI
	Restaurant() RestaurantI
}

type MenuI interface {
	Create(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error)
	Get(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error)
	Update(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error)
	Delete(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error)
	GetAll(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error)
}


type ReservationI interface{
	Create(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error)
	Get(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationResponse, error)
	Update(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationResponse, error)
	Delete(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error)
	GetAll(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error)
	ListReservation(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error)
	CheckAvailability(ctx context.Context, req *reservation.CheckAvailabilityRequest) (*reservation.CheckAvailabilityResponse, error)
	FoodList(ctx context.Context, req *reservation.OrderFoodListReq) (*reservation.OrderFoodListRes, error)
	OrderFood(ctx context.Context, req *reservation.OrderFoodReq) (*reservation.OrderFoodRes, error	)
	IsValidReservation(ctx context.Context, req *reservation.IsValidReq) (*reservation.IsValidRes, error)
	OrderBill(ctx context.Context, req *reservation.OrderBillReq) (*reservation.OrderBillRes, error)
}

type RestaurantI interface{
	Create(ctx context.Context, req *restaurant.CreateRestaurantRequest) (*restaurant.CreateRestaurantResponse, error)
	Get(ctx context.Context, req *restaurant.GetRestaurantRequest) (*restaurant.GetRestaurantResponse, error)
	Update(ctx context.Context, req *restaurant.UpdateRestaurantRequest) (*restaurant.UpdateRestaurantResponse, error)
	Delete(ctx context.Context, req *restaurant.DeleteRestaurantRequest) (*restaurant.DeleteRestaurantResponse, error)
	GetAll(ctx context.Context, req *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error)

}