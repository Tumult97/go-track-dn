package responses

import "time"

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token,omitempty"`
	Message   string    `json:"message"`
}

