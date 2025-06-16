package responses

import "time"

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Message      string    `json:"message"`
}

// RefreshTokenResponse represents the response for token refresh operations
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"` // Optional: new refresh token for rotation
	Message      string `json:"message"`
}
