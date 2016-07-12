package config

// NewConfig will read in the config.json file and return a struct representation
func NewConfig() Config {
	return Config{
		APIEndpoint: "the-api-endpoint",
		AuthToken:   "some-RSA-token",
		SpaceGUID:   "some-space-guid",
	}
}

// Config is the struct representation of the config.json file
type Config struct {
	APIEndpoint string
	AuthToken   string
	SpaceGUID   string
}
