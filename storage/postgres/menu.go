package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Diyor/Project/Restaurant/genproto/menu"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Menu struct {
	db *sql.DB
}

func NewMenu(db *sql.DB) *Menu {
	return &Menu{db: db}
}

func (m *Menu) Create(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error) {
	id := uuid.New().String()
	query := `
	insert into menus
	(id, restaurant_id, name,description, price)
	values
	($1, $2, $3, $4, $5)
	RETURNING id, restaurant_id, name, description, price
	`
	var res menu.Menu
	err := m.db.QueryRowContext(ctx, query, id, req.Menu.RestaurantId, req.Menu.Name, req.Menu.Description, req.Menu.Price).
		Scan(res.Id, res.RestaurantId, res.Name, res.Description, res.Price)
	if err != nil {
		return nil, err
	}
	return &menu.CreateMenuResponse{Menu: &res}, nil
}

func (m *Menu) Get(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error) {
	res := menu.Menu{}
	var createdAt time.Time
	var deleteddAt time.Time
	query := `
	select id, restaurant_id, name, description, price
	from menus
	where id = $1
	`
	err := m.db.QueryRowContext(ctx, query, req.Id).
		Scan(&res.Id, &res.RestaurantId, &res.Name, &res.Description, &res.Price, &createdAt, &deleteddAt)
	if err != nil {
		return nil, err
	}
	return &menu.GetMenuResponse{Menu: &res}, nil
}

func (m *Menu) Update(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error) {
	query := `
	update menus
	set name = $1, description = $2, price = $3
	where id = $4
	RETURNING id, restaurant_id, name, description, price
	`
	var res menu.Menu
	err := m.db.QueryRowContext(ctx, query, req.Menu.Name, req.Menu.Description, req.Menu.Price, req.Menu.Id).
		Scan(&res.Id, &res.RestaurantId, &res.Name, &res.Description, &res.Price)
	if err != nil {
		return nil, err
	}
	return &menu.UpdateMenuResponse{Menu: &res}, nil
}

func (m *Menu) Delete(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error) {
	query := `
	delete from menus
	where id = $1
	`
	_, err := m.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	return &menu.DeleteMenuResponse{}, nil
}

func (m *Menu) GetAll(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			restaurant_id, 
			name,
			description,
			price,
			created_at,
			updated_at,
			deleted_at
		FROM 
			menus
		WHERE 
			deleted_at = 0
	`
	filter := ""

	if req.RestaurantId != "" {
		filter += fmt.Sprintf(" AND restaurant_id = $%d", count)
		args = append(args, req.RestaurantId)
		count++
	}

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count) // ILIKE for case-insensitive search
		args = append(args, "%"+req.Name+"%")               // Add wildcards for partial match
		count++
	}

	if req.Description != "" {
		filter += fmt.Sprintf(" AND description ILIKE $%d", count)
		args = append(args, "%"+req.Description+"%")
		count++
	}

	if req.MinPrice != 0 {
		filter += fmt.Sprintf(" AND price >= $%d", count)
		args = append(args, req.MinPrice)
		count++
	}

	if req.MaxPrice != 0 {
		filter += fmt.Sprintf(" AND price <= $%d", count)
		args = append(args, req.MaxPrice)
		count++
	}

	query += filter

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching menus from the database")
		return nil, err
	}
	defer rows.Close()

	var menus []*menu.Menu
	for rows.Next() {
		var (
			createdAt time.Time
			updatedAt time.Time
			m         menu.Menu
		)

		err := rows.Scan(
			&m.Id,
			&m.RestaurantId,
			&m.Name,
			&m.Description,
			&m.Price,
			&createdAt,
			&updatedAt,
			&m.DeletedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning menu row")
			return nil, err
		}
		m.CreatedAt = createdAt.Format(time.RFC3339)
		m.UpdatedAt = updatedAt.Format(time.RFC3339)

		menus = append(menus, &m)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over menu rows")
		return nil, err
	}

	return &menu.GetAllMenusResponse{Menus: menus}, nil
}
