package plaid

import (
	"encoding/json"
	"fmt"
)

type authResponse struct {
	Accounts  []Account `json:"accounts"`
	Numbers   []Number  `json:"numbers"`
	Item      `json:"item"`
	RequestID string `json:"request_id"`
}

// Auth fetches data from a user already in the auth product
func (i Item) Auth(clnt Client) (Auth, error) {
	bts, err := get(fmt.Sprintf("%v/auth/get", clnt.envURL), clnt, i)
	if err != nil {
		return Auth{}, err
	}
	ar := authResponse{}
	err = json.Unmarshal(bts, &ar)
	return Auth{
		Accounts: ar.Accounts,
		Numbers:  ar.Numbers,
	}, err
}
