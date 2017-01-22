package plaid

import "encoding/json"

// Auth grants access to the Auth product
type Auth interface {
	AddMFA() ([]Account, string, error)
	GetData()
	AddUser(string, string, string, string) ([]Account, string, error)
	UpdateUser()
	DeleteUser()
}

// auth sets up an Auth service client
type auth struct {
	Client
	remote string
}

const (
	authURL = "/auth"
)

// UseAuth configures an auth product for use
func UseAuth(clnt Client) Auth {
	return &auth{
		remote: clnt.envURL + authURL,
		Client: clnt,
	}
}

// Account ...
type Account struct {
	ID              string `json:"_id"`
	Item            string `json:"_item"`
	User            string `json:"_user"`
	InstitutionType string `json:"institution_type"`
	Type            string `json:"depository"`
	Balance         struct {
		Available float64 `json:"available"`
		Current   float64 `json:"current"`
	} `json:"balance"`
	Meta struct {
		Name   string `json:"name"`
		Number string `json:"nuber"`
	} `json:"meta"`
	Numbers struct {
		Routing     string `json:"routing"`
		Account     string `json:"account"`
		WireRouting string `json:"wireRouting"`
	} `json:"numbers"`
}

// authRequest is used to send an auth login request to Plaid
type authRequest struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Pin      string `json:"pin,omitempty"`
}

// accountRes is used to unmarshal accounts
type accountsRes struct {
	Accounts    []Account `json:"accounts"`
	AccessToken string    `json:"access_token"`
}

// AddMFA ...
func (a *auth) AddMFA() ([]Account, string, error) {
	return nil, "", nil
}

// GetData ...
func (a *auth) GetData() {}

// AddUser ...
func (a *auth) AddUser(username, password, institution, pin string) ([]Account, string, error) {
	bts, err := json.Marshal(authRequest{
		ClientID: a.clientID,
		Secret:   a.clientSecret,
		Type:     institution,
		Username: username,
		Password: password,
		Pin:      pin,
	})
	if err != nil {
		return nil, "", err
	}
	// TODO: add post
	ar := accountsRes{}
	if err := json.Unmarshal(bts, &ar); err != nil {
		return nil, "", err
	}
	return ar.Accounts, ar.AccessToken, nil
}

// UpdateUser ...
func (a *auth) UpdateUser() {}

// DeleteUser ...
func (a *auth) DeleteUser() {}
