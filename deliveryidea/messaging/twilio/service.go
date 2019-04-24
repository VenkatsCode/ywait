package twilio

import (
	"github.com/Bimde/httputils"
	log "github.com/sirupsen/logrus"
)

func CreateClient(config *Config) *Client {
	return &Client{Config: config, http: httputils.New(config.AccountID, config.AuthToken)}
}

func (c *Client) SendMessage(msg *MessageRequest) error {
	output := map[string]interface{}{}
	c.fillDefaults(msg)
	err := c.http.Post(nil, c.getURL(), msg, output)
	if err != nil {
		log.WithField("message_body", msg).Error("Error sending message")
		return err
	}
	log.WithField("message_body", msg).Debug("Sent message")
	return nil
}

func (c *Client) fillDefaults(msg *MessageRequest) {
	if msg.From == "" {
		msg.From = c.Config.From
	}
}

func (c *Client) getURL() string {
	return "https://api.twilio.com/2010-04-01/Accounts/" + c.Config.AccountID + "Messages.json"
}
