package domain

// AuthRequest represents the request to get a token
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse represents the response containing the token
type AuthResponse struct {
	Token string `json:"token"`
}
