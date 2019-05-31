// Package twilio provides and interface for the Twilio REST API
package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Bimde/httputils"
	log "github.com/sirupsen/logrus"
)

func CreateClient(config *Config) *Client {
	return &Client{Config: config, http: httputils.New(config.AccountID, config.AuthToken)}
}

func (c *Client) SendMessage(msg *MessageRequest) error {
	output := map[string]interface{}{}
	log.WithField("url", c.getURL()).Info("Message sending, ", msg)
	// c.fillDefaults(msg)
	msg.From = "+16479303390"
	// err := c.http.Post(nil, c.getURL(), msg, output)
	msgData := url.Values{}
	msgData.Set("To", msg.To)
	msgData.Set("From", msg.From)
	msgData.Set("Body", msg.Body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", c.getURL(), &msgDataReader)
	req.SetBasicAuth(c.Config.AccountID, c.Config.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// if err != nil {
	// 	log.WithField("message_body", msg).Error("Error sending message")
	// 	return err
	// }

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}

	log.WithField("message_body", msg).WithField("output", output).Debug("Sent message")
	return nil
}

func (c *Client) fillDefaults(msg *MessageRequest) {
	if msg.From == "" {
		msg.From = c.Config.From

	}
}

func (c *Client) getURL() string {
	return "https://api.twilio.com/2010-04-01/Accounts/" + c.Config.AccountID + "/Messages.json"
}
