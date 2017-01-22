package plaid

import (
	"bytes"
	"encoding/json"
)

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

// AddUser adds an auth user to plaid and returns a slice of accounts, an access
// token or an error if these fail. If the user requires MFA then nil will be
// passed instead of accounts, an access token will be passed, and an
// MfaRequired error is returned
func (a *auth) AddUser(username, password, institution, pin string) ([]Account, string, error) {
	// build request
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
	// send request and parse errors
	res, err := post(a.remote, bytes.NewBuffer(bts))
	if err != nil && err == MfaRequired {
		// return mfa required and pull the access token out of the response
		ar := accountsRes{}
		if err = json.Unmarshal(res, &ar); err != nil {
			return nil, "", err
		}
		return nil, ar.AccessToken, MfaRequired
	}
	if err != nil {
		return nil, "", err
	}
	// unmarshal and return the accessed accounts and the access token
	ar := accountsRes{}
	if err := json.Unmarshal(res, &ar); err != nil {
		return nil, "", err
	}
	return ar.Accounts, ar.AccessToken, nil
}

// UpdateUser ...
func (a *auth) UpdateUser() {}

// DeleteUser ...
func (a *auth) DeleteUser() {}
