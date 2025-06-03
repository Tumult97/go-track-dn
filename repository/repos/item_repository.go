package repos

import (
	"context"
	"database/sql"
	"fmt"

	entities "tracker/models/entities"
	database "tracker/repository"
)

type IItemRepository interface {
	Create(item entities.Item) (int, error)
	GetAll() (*[]entities.Item, error)
	GetById(id int) (*entities.Item, error)
}

type itemRepository struct{}

func NewItemRepository() IItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) Create(item entities.Item) (int, error) {
	query := `
		INSERT INTO items (name, description, quantity, price, is_per_item)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2, @p3, @p4, @p5)
	`

	var id int

	err := database.DB.QueryRowContext(context.Background(), query,
		sql.Named("p1", item.Name),
		sql.Named("p2", item.Description),
		sql.Named("p3", item.Quantity),
		sql.Named("p4", item.Price),
		sql.Named("p5", item.IsPerItem),
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert item: %w", err)
	}

	return id, nil
}

func (r *itemRepository) GetAll() (*[]entities.Item, error) {
	query := `SELECT id, name, description, quantity, price, is_per_item FROM items`
	rows, err := database.DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query items: %w", err)
	}
	defer rows.Close()

	var items []entities.Item
	for rows.Next() {
		var item entities.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.Price, &item.IsPerItem)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	return &items, nil
}

func (r *itemRepository) GetById(id int) (*entities.Item, error) {
	query := `SELECT id, name, description, quantity, price, is_per_item FROM items WHERE id = @p1`
	row := database.DB.QueryRowContext(context.Background(), query, sql.Named("p1", id))

	var item entities.Item
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.Price, &item.IsPerItem)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item by id: %w", err)
	}

	return &item, nil
}
