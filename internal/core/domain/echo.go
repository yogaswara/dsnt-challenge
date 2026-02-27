package domain

import "encoding/json"

// EchoRequest represents the incoming request for echo
type EchoRequest = json.RawMessage

// EchoResponse represents the output for echo
type EchoResponse = json.RawMessage
