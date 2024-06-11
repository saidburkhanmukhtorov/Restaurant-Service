package storage

import (
	"context"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/genproto/menu"
)

type StorageI interface {
	Menu() MenuI
	Reservation() ReservationI
	Restaurant() RestaurantI
}

// MenuI defines methods for managing menu items.
type MenuI interface {
	Create(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error)
	GetById(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error)
	Update(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error)
	Delete(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error)
	// Add other menu-related methods as needed
	GetAll(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error)
}

type ReservationI interface{}
type RestaurantI interface{}
