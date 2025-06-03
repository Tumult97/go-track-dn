package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	entities "tracker/models/entities"
	database "tracker/repository"
)

// IUserRepository interface defines methods for user management
type IUserRepository interface {
	CreateUser(user entities.User) (int, error)
	GetUserByEmail(email string) (*entities.User, error)
	ValidateCredentials(email, password string) (*entities.User, error)
}

type userRepository struct{}

// NewUserRepository creates a new instance of the user repository
func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(user entities.User) (int, error) {
	// Hash the password before storing
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	query := `
		INSERT INTO users (first_name, last_name, email, password, created_at)
		OUTPUT INSERTED.id
		VALUES (@firstName, @lastName, @email, @password, @createdAt)
	`

	var id int
	err = database.DB.QueryRowContext(context.Background(), query,
		sql.Named("firstName", user.FirstName),
		sql.Named("lastName", user.LastName),
		sql.Named("email", user.Email),
		sql.Named("password", hashedPassword),
		sql.Named("createdAt", time.Now()),
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

// GetUserByEmail retrieves a user by their email address
func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = @userEmail`
	
	row := database.DB.QueryRowContext(context.Background(), query, sql.Named("userEmail", email))

	var user entities.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with email %s", email)
		}
		return nil, fmt.Errorf("failed to fetch user by email: %w", err)
	}

	return &user, nil
}

// ValidateCredentials checks if the provided email and password are valid
func (r *userRepository) ValidateCredentials(email, password string) (*entities.User, error) {
	// Get user by email
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the stored hash
	err = comparePasswords(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Password is correct, return user without password
	user.Password = ""
	return user, nil
}

// Helper function to hash passwords
func hashPassword(password string) (string, error) {
	// Generate a bcrypt hash with cost factor 12
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Helper function to compare passwords
func comparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

