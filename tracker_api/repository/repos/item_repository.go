package repos

import (
	"context"
	"database/sql"
	"fmt"

	entities "tracker/models/entities"
	database "tracker/repository"
)

type IItemRepository interface {
	Create(item *entities.Item) (*entities.Item, error)
	Update(item *entities.Item) (*entities.Item, error)
	GetAll(userId int) (*[]entities.Item, error)
	GetById(id int, userId int) (*entities.Item, error)
	DeleteById(id int, userId int) (int, error)
}

type itemRepository struct{}

func NewItemRepository() IItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) Create(item *entities.Item) (*entities.Item, error) {
	query := `
		INSERT INTO items (name, description, quantity, price, is_per_item, user_id)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6)
	`

	var id int

	err := database.DB.QueryRowContext(context.Background(), query,
		sql.Named("p1", item.Name),
		sql.Named("p2", item.Description),
		sql.Named("p3", item.Quantity),
		sql.Named("p4", item.Price),
		sql.Named("p5", item.IsPerItem),
		sql.Named("p6", item.UserId),
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %w", err)
	}

	item.Id = id

	return item, nil
}

func (r *itemRepository) Update(item *entities.Item) (*entities.Item, error) {
	fmt.Println("Inside Update method")
	query := `
		UPDATE dbo.items
			SET name = @name,
			description = @description,
			quantity = @quantity,
			price = @price,
			is_per_item = @is_per_item
			user_id = @user_id
		OUTPUT INSERTED.id
		WHERE id = @id
	`

	var id int

	err := database.DB.QueryRowContext(context.Background(), query,
		sql.Named("id", item.Id),
		sql.Named("name", item.Name),
		sql.Named("description", item.Description),
		sql.Named("quantity", item.Quantity),
		sql.Named("price", item.Price),
		sql.Named("is_per_item", item.IsPerItem),
		sql.Named("user_id", item.UserId),
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	return item, nil
}

func (r *itemRepository) GetAll(userId int) (*[]entities.Item, error) {
	query := `SELECT id, name, description, quantity, price, is_per_item, user_id FROM items WHERE user_id = @userId`
	rows, err := database.DB.QueryContext(context.Background(), query, sql.Named("userId", userId))
	if err != nil {
		return nil, fmt.Errorf("failed to query items: %w", err)
	}
	defer rows.Close()

	var items []entities.Item
	for rows.Next() {
		var item entities.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.Price, &item.IsPerItem, &item.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	return &items, nil
}

func (r *itemRepository) GetById(id int, userId int) (*entities.Item, error) {
	query := `SELECT TOP(1) id, name, description, quantity, price, is_per_item, user_id FROM items WHERE id = @id AND user_id = @userId`
	row := database.DB.QueryRowContext(context.Background(), query, sql.Named("id", id), sql.Named("userId", userId))

	var item entities.Item
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.Price, &item.IsPerItem)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item by id: %w", err)
	}

	return &item, nil
}

func (r *itemRepository) DeleteById(id int, userId int) (int, error) {
	query := "DELETE FROM items WHERE id = @id AND user_id = @userId"
	err := database.DB.QueryRowContext(context.Background(), query, sql.Named("id", id), sql.Named("userId", userId)).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to delete item: %w", err)
	}

	return id, nil
}
