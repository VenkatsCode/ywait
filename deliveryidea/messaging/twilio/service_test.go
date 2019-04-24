package twilio

import (
	"testing"
)

var (
	client *Client
)

func TestCreateCart(t *testing.T) {
	cartsSetup()
	defer cartsShutdown()

}

func cartsSetup() {
	client = CreateClient(&Config{AccountID})
}

func cartsShutdown() {

}
