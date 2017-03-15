package plaid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Option allows users to set configurable options for the Plaid client
type Option func(*Client) error

// Client stores client info for Plaid and is passed to utilize Plaid products
type Client struct {
	envURL       string
	clientID     string
	clientSecret string
}

const (
	// Development points to the Plaid UAT environment
	Development = "https://development.plaid.com/"

	// Production points to the Plaid Production environment
	Production = "https://production.plaid.com/"

	// Sandbox points to the Plaid Sandbox environment
	Sandbox = "https://sandbox.plaid.com"
)

// Configure sets up a plaid client and returns interfaces that
// can be used to request the various products
func Configure(clientID, clientSecret string, ops ...Option) (Client, error) {
	clnt := Client{
		envURL:       Development,
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

/*** Options ***/

// SetEnvironment sets the remote URL to which the client will send requests
func SetEnvironment(url string) Option {
	return func(c *Client) error {
		if url != Development && url != Production && url != Sandbox {
			return fmt.Errorf("Must select either %v, %v, or %v for url", Development, Sandbox, Production)
		}
		c.envURL = url
		return nil
	}
}

// SetWebhooks sets up the plaid client to handle webhooks
func SetWebhooks() Option {
	return func(c *Client) error {
		return nil
	}
}

/*** Abstract http methods ***/

// getRequest ...
type getRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

func get(url string, clnt Client, i Item) ([]byte, error) {
	bts, err := json.Marshal(getRequest{
		ClientID:    clnt.clientID,
		Secret:      clnt.clientSecret,
		AccessToken: i.AccessToken,
	})
	if err != nil {
		return nil, err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	bts, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, formatError(bts)
	}
	return bts, nil
}
