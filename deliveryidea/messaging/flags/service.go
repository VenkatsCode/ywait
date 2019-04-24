// Package config provides an abstraction for application configuration.
// Configuration could come from the environment, config files, or be directly
// set programatically.
package flags

var config map[string]interface{}

const (
	TwilioAccountID  = "account_id"
	TwilioAuthToken  = "auth_token"
	TwilioFromNumber = "twilio_from_number"
)

func init() {
	config = map[string]interface{}{}

	// TODO get these from environment
	config[TwilioAccountID] = "ACc5eda94fc61ce2ce2f955e1ec73f15e8"
	config[TwilioAuthToken] = "7062d893bda6fdca3f34dadafe043140"
	config[TwilioFromNumber] = "+16479303390"
}

// Get returns a property from the configuration with the given key
func Get(property string) interface{} {
	return config[property]
}

// GetString returns a property from the configuration type-asserted to a string
func GetString(property string) string {
	return config[property].(string)
}

// Set allows overriding of the environment configuration programatically
func Set(key string, value interface{}) {
	config[key] = value
}
