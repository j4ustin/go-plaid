package plaid

import (
	"bytes"
	"encoding/json"
)

// Auth grants access to the Auth product
type Auth interface {
	MfaStep(string, string) ([]Account, string, MFA, error)
	GetData(string) ([]Account, MFA, error)
	AddUser(string, string, string, string) ([]Account, string, MFA, error)
	UpdateUser(string, string, string) ([]Account, string, MFA, error)
	DeleteUser(string) error
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

// Account contains bank info after a successful auth
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

// MFA is returned with instructions to handle an MFA option
type MFA struct {
	Type        string            `json:"type"`
	Mfa         map[string]string `json:"mfa"`
	AccessToken string            `json:"access_token"`
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

// mfaRequest is used to send a response to an MFA step
type mfaRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	Mfa         string `json:"mfa"`
}

// getRequest ...
type getRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

// updateRequest ...
type updateRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

type deleteRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

// accountRes is used to unmarshal accounts
type accountsRes struct {
	Accounts    []Account `json:"accounts"`
	AccessToken string    `json:"access_token"`
}

// MfaStep is used to request a new MFA type
func (a *auth) MfaStep(accessToken, mfaAnswer string) ([]Account, string, MFA, error) {
	bts, err := json.Marshal(mfaRequest{
		ClientID:    a.clientID,
		Secret:      a.clientSecret,
		AccessToken: accessToken,
		Mfa:         mfaAnswer,
	})
	if err != nil {
		return nil, "", MFA{}, err
	}
	return handleAuthResponse(post(a.remote+"/step", bytes.NewBuffer(bts)))
}

// GetData fetches data from a user already in the auth product
func (a *auth) GetData(accessToken string) ([]Account, MFA, error) {
	bts, err := json.Marshal(getRequest{
		ClientID:    a.clientID,
		Secret:      a.clientSecret,
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, MFA{}, err
	}
	accts, _, mfa, err := handleAuthResponse(post(a.remote+"/get", bytes.NewBuffer(bts)))
	return accts, mfa, err
}

// AddUser adds an auth user to plaid and returns a slice of accounts, an access
// token or an error if these fail. If the user requires MFA then nil will be
// passed instead of accounts, an access token will be passed, and an
// MfaRequired error is returned
func (a *auth) AddUser(username, password, institution, pin string) ([]Account, string, MFA, error) {
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
		return nil, "", MFA{}, err
	}
	// send request and parse errors
	return handleAuthResponse(post(a.remote, bytes.NewBuffer(bts)))
}

// UpdateUser is used to update credentials to a user's bank
func (a *auth) UpdateUser(username, password, accessToken string) ([]Account, string, MFA, error) {
	bts, err := json.Marshal(updateRequest{
		ClientID:    a.clientID,
		Secret:      a.clientSecret,
		Username:    username,
		Password:    password,
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, "", MFA{}, err
	}
	return handleAuthResponse(patch(a.remote, bytes.NewBuffer(bts)))
}

// DeleteUser ...
func (a *auth) DeleteUser(accessToken string) error {
	bts, err := json.Marshal(deleteRequest{
		ClientID:    a.clientID,
		Secret:      a.clientSecret,
		AccessToken: accessToken,
	})
	if err != nil {
		return err
	}
	_, err = delete(a.remote, bytes.NewBuffer(bts))
	return err
}

func handleAuthResponse(res []byte, err error) ([]Account, string, MFA, error) {
	if err != nil && err == ErrMfaRequired {
		// return mfa required and pull the access token out of the response
		mfa := MFA{}
		if err = json.Unmarshal(res, &mfa); err != nil {
			return nil, "", MFA{}, err
		}
		return nil, mfa.AccessToken, mfa, ErrMfaRequired
	}
	if err != nil {
		return nil, "", MFA{}, err
	}
	// unmarshal and return the accessed accounts and the access token
	ar := accountsRes{}
	if err := json.Unmarshal(res, &ar); err != nil {
		return nil, "", MFA{}, err
	}
	return ar.Accounts, ar.AccessToken, MFA{}, nil
}
