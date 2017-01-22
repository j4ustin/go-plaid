package plaid

import "fmt"

// Options allows users to set configurable options for the Plaid client
type Options func(*Client) error

// Client stores client info for Plaid and is passed to utilize Plaid products
type Client struct {
	envURL       string
	clientID     string
	clientSecret string
}

const (
	// Uat points to the Plaid UAT environment
	Uat = "https://tartan.plaid.com/"

	// Production points to the plaid Production environment
	Production = "https://api.plaid.com/"
)

// Configure sets up a plaid client and returns interfaces that
// can be used to request the various products
func Configure(clientID, clientSecret string, ops ...Options) (Client, error) {
	clnt := Client{
		envURL:       Uat,
		clientSecret: clientSecret,
		clientID:     clientID,
	}
	for _, op := range ops {
		if err := op(&clnt); err != nil {
			return Client{}, err
		}
	}
	return clnt, nil
}

// SetEnvironment sets the remote URL to which the client will send requests
func SetEnvironment(url string) Options {
	return func(c *Client) error {
		if url != Uat && url != Production {
			return fmt.Errorf("Must select either %v or %v for url", Uat, Production)
		}
		c.envURL = url
		return nil
	}
}
