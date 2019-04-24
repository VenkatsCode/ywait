package twilio

import "github.com/Bimde/httputils"

// Message represents the body of the message data sent to Twilio
type MessageRequest struct {
	To   string `json:"To"`
	From string `json:"From"`
	Body string `json:"Body"`
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
