package repos

import (
	"context"
	"database/sql"
	"fmt"

	entities "tracker/models/entities"
	database "tracker/repository"
)

type ILocationRepository interface {
	Create(location entities.Location) (int, error)
	Update(location entities.Location) (int, error)
	GetById(id int, userId int) (*entities.Location, error)
	GetAll(userId int) (*[]entities.Location, error)
	DeleteById(id int, userId int) (int, error)
}

type locationRepository struct{}

func NewLocationRepository() ILocationRepository {
	return &locationRepository{}
}

func (repo *locationRepository) Create(location entities.Location) (int, error) {
	query := `
	    INSERT INTO dbo.location
		(
			name,
			description,
			address_home,
			address_street,
			address_suburb,
			address_city,
			user_id
		)
		OUTPUT INSERTED.id
		VALUES
		(   @name,   		-- name - varchar(50)
			@description, 	-- description - varchar(256)
			@home, 			-- address_home - varchar(256)
			@street, 		-- address_street - varchar(256)
			@suburb, 		-- address_suburb - varchar(256)
			@city, 			-- address_city - varchar(256)
			@userId     	-- user_id - int
			);
	`

	var id int

	err := database.DB.QueryRowContext(context.Background(), query,
		sql.Named("name", location.Name),
		sql.Named("description", location.Description),
		sql.Named("home", location.AddressHome),
		sql.Named("street", location.AddressStreet),
		sql.Named("suburb", location.AddressSuburb),
		sql.Named("city", location.AddressCity),
		sql.Named("userId", location.UserID),
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert location: %w", err)
	}

	return id, nil
}

func (repo *locationRepository) Update(location entities.Location) (int, error) {
	query := `
		UPDATE dbo.location
		SET name = @name,
			description = @description,
			address_home = @home,
			address_street = @street,
			address_suburb = @suburb,
			address_city = @city
		WHERE id = @id AND user_id = @userId
		OUTPUT INSERTED.id;
	`

	var id int

	err := database.DB.QueryRowContext(context.Background(), query,
		sql.Named("name", location.Name),
		sql.Named("description", location.Description),
		sql.Named("home", location.AddressHome),
		sql.Named("street", location.AddressStreet),
		sql.Named("suburb", location.AddressSuburb),
		sql.Named("city", location.AddressCity),
		sql.Named("id", location.ID),
		sql.Named("userId", location.UserID),
	).Scan(&id)

	if err != nil || id == 0 {
		return 0, fmt.Errorf("failed to update location: %w", err)
	}

	return id, nil
}

func (repo *locationRepository) GetById(id int, userId int) (*entities.Location, error) {
	query := `SELECT TOP(1) * FROM location where id = @id AND user_id = @userId;`

	var location entities.Location

	row := database.DB.QueryRowContext(context.Background(), query, sql.Named("id", id), sql.Named("userId", userId))

	err := row.Scan(&location.ID,
		&location.Name,
		&location.Description,
		&location.AddressHome,
		&location.AddressStreet,
		&location.AddressSuburb,
		&location.AddressCity,
		&location.UserID)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch location by id: %w", err)
	}

	return &location, nil
}

func (repo *locationRepository) GetAll(userId int) (*[]entities.Location, error) {
	query := `SELECT * FROM location where user_id = @userId;`

	var locations []entities.Location

	rows, err := database.DB.QueryContext(context.Background(), query, sql.Named("userId", userId))

	if err != nil {
		return nil, fmt.Errorf("failed to query items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var location entities.Location

		err := rows.Scan(
			&location.ID,
			&location.Name,
			&location.Description,
			&location.AddressHome,
			&location.AddressStreet,
			&location.AddressSuburb,
			&location.AddressCity,
			&location.UserID)

		if err != nil {
			return nil, fmt.Errorf("failed to scan location: %w", err)
		}
		locations = append(locations, location)
	}

	return &locations, nil
}

func (repo *locationRepository) DeleteById(id int, userId int) (int, error) {
	query := `UPDATE dbo.items
			SET location_id = NULL
			WHERE location_id = @locationId
			AND user_id = @userId;

			DELETE FROM dbo.location
			OUTPUT DELETED.id
			WHERE id = @locationId
			AND user_id = @userId;`

	err := database.DB.QueryRowContext(context.Background(), query,
		sql.Named("locationId", id),
		sql.Named("userId", userId)).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error deleting location %w", err)
	}

	return id, nil
}
