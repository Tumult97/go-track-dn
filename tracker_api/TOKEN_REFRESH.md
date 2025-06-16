# Token Refresh System

This application implements a secure token refresh system using JWT (JSON Web Tokens) with separate access and refresh tokens.

## Overview

- **Access Tokens**: Short-lived tokens (default: 15 minutes) used for API authentication
- **Refresh Tokens**: Long-lived tokens (default: 7 days) used to generate new access tokens
- **Token Rotation**: New refresh tokens are issued with each refresh to enhance security

## Configuration

Set the following environment variables:

```bash
# Required
JWT_SECRET=your-super-secret-jwt-key-here
TOKEN_ISSUER=tracker-api

# Optional - Token expiration settings
ACCESS_TOKEN_EXPIRATION_MINUTES=15  # Default: 15 minutes
REFRESH_TOKEN_EXPIRATION_DAYS=7     # Default: 7 days
```

## API Endpoints

### Login
```
POST /auth/login
```

**Request:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "user@example.com",
  "created_at": "2023-01-01T00:00:00Z",
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "token": "eyJhbGciOiJIUzI1NiIs...",  // For backward compatibility
  "message": "Login successful"
}
```

### Refresh Token
```
POST /auth/refresh
```

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",  // New refresh token
  "message": "Token refreshed successfully"
}
```

## Usage Flow

1. **Initial Login**: User logs in and receives both access and refresh tokens
2. **API Requests**: Use the access token in the Authorization header: `Bearer <access_token>`
3. **Token Expiry**: When access token expires (15 minutes), use refresh token to get new tokens
4. **Token Rotation**: Each refresh generates a new refresh token for enhanced security
5. **Logout**: Discard both tokens on the client side

## Security Features

- **Short-lived Access Tokens**: Minimize exposure window if compromised
- **Token Type Validation**: Access and refresh tokens are validated for their specific use
- **Token Rotation**: New refresh tokens issued with each refresh
- **Separate Claims**: Different token structures for access vs refresh tokens

## Client Implementation Example

```javascript
class TokenManager {
  constructor() {
    this.accessToken = localStorage.getItem('access_token');
    this.refreshToken = localStorage.getItem('refresh_token');
  }

  async makeAuthenticatedRequest(url, options = {}) {
    try {
      // Try with current access token
      const response = await fetch(url, {
        ...options,
        headers: {
          ...options.headers,
          'Authorization': `Bearer ${this.accessToken}`
        }
      });

      if (response.status === 401) {
        // Access token expired, try to refresh
        await this.refreshTokens();
        
        // Retry the original request
        return await fetch(url, {
          ...options,
          headers: {
            ...options.headers,
            'Authorization': `Bearer ${this.accessToken}`
          }
        });
      }

      return response;
    } catch (error) {
      console.error('Request failed:', error);
      throw error;
    }
  }

  async refreshTokens() {
    const response = await fetch('/auth/refresh', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        refresh_token: this.refreshToken
      })
    });

    if (!response.ok) {
      // Refresh failed, redirect to login
      this.logout();
      throw new Error('Refresh failed');
    }

    const data = await response.json();
    this.accessToken = data.access_token;
    this.refreshToken = data.refresh_token;
    
    localStorage.setItem('access_token', this.accessToken);
    localStorage.setItem('refresh_token', this.refreshToken);
  }

  logout() {
    this.accessToken = null;
    this.refreshToken = null;
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    window.location.href = '/login';
  }
}
```

## Migration Notes

For existing applications:

1. The `token` field in login response is maintained for backward compatibility
2. Existing middleware continues to work with access tokens
3. Update clients to use `access_token` and `refresh_token` fields
4. Implement refresh logic before access tokens expire

