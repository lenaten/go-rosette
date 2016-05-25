package rosette

import "github.com/segmentio/go-env"

func client() *Client {
	apiKey := env.MustGet("ROSETTE_API_KEY")
	return New(NewConfig(apiKey))
}
