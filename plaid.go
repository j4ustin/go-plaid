package plaid

// Options allows users to set configurable options for the Plaid client
type Options func(client)

type client struct {
	envURL       string
	clientID     string
	clientSecret string
}

// Configure sets up a plaid client and returns interfaces that
// can be used to request the various products
func Configure(clientID, clientSecret string, ops ...Options) {}
