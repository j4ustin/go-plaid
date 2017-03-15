package plaid

import (
	"encoding/json"
	"fmt"
)

type accountResponse struct {
	Item      `json:"item"`
	RequestID string    `json:"request_id"`
	Accounts  []Account `json:"accounts"`
}

// Accounts returns accounts associated with the access token
func (i Item) Accounts(clnt Client) ([]Account, error) {
	bts, err := get(fmt.Sprintf("%v/accounts/get", clnt.envURL), clnt, i)
	if err != nil {
		return nil, err
	}
	ar := accountResponse{}
	err = json.Unmarshal(bts, &ar)
	return ar.Accounts, err
}
