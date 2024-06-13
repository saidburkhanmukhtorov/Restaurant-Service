package service

import (
	"context"

	"github.com/Diyor/Project/Restaurant/genproto/menu"
	"github.com/Diyor/Project/Restaurant/storage"
)

type MenuService struct {
	s storage.StorageI
	menu.UnimplementedMenuServiceServer
}

func NewMenuService(s storage.StorageI) *MenuService {
	return &MenuService{s: s}
}

func (s *MenuService) CreateMenu(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error) {
	res, err := s.s.Menu().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil

}
func (s *MenuService) GetMenu(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error) {
	res, err := s.s.Menu().Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error) {
	res, err := s.s.Menu().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MenuService) DeleteMenu(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error) {
	res, err := s.s.Menu().Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MenuService) GetAllMenus(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error) {
	res, err := s.s.Menu().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

