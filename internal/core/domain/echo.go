package domain

// EchoRequest represents the incoming request for echo
type EchoRequest struct {
	Message string `json:"message"`
}

// EchoResponse represents the output for echo
type EchoResponse struct {
	Message string `json:"message"`
}
