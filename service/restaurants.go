package service


import (
	"context"

	"github.com/Diyor/Project/Restaurant/genproto/restaurant"
	"github.com/Diyor/Project/Restaurant/storage"
)

type RestaurantService struct {
	s storage.StorageI
	restaurant.UnimplementedRestaurantServiceServer
}

func NewRestaurantService(s storage.StorageI) *RestaurantService {
	return &RestaurantService{s: s}
}

func (s *RestaurantService) CreateRestaurant(ctx context.Context, req *restaurant.CreateRestaurantRequest) (*restaurant.CreateRestaurantResponse, error) {
	res, err := s.s.Restaurant().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil

}
func (s *RestaurantService) GetRestaurant(ctx context.Context, req *restaurant.GetRestaurantRequest) (*restaurant.GetRestaurantResponse, error) {
	res, err := s.s.Restaurant().Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RestaurantService) UpdateRestaurant(ctx context.Context, req *restaurant.UpdateRestaurantRequest) (*restaurant.UpdateRestaurantResponse, error) {
	res, err := s.s.Restaurant().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RestaurantService) DeleteRestaurant(ctx context.Context, req *restaurant.DeleteRestaurantRequest) (*restaurant.DeleteRestaurantResponse, error) {
	res, err := s.s.Restaurant().Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *RestaurantService) ListRestaurants(ctx context.Context, req *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error) {
	res, err := s.s.Restaurant().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}