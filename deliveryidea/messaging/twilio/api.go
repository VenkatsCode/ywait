package twilio

import "github.com/Bimde/httputils"

// Message represents the body of the message data sent to Twilio
type MessageRequest struct {
	To   string `json:"to"`
	From string `json:"from"`
	Body string `json:"body"`
}

// Client contains Twilio configuration
type Client struct {
	Config *Config
	http   *httputils.Client
}

// Config contains Twillio connection info
type Config struct {
	AccountID string
	AuthToken string
	From      string
}
