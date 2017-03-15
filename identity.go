package plaid

import (
	"encoding/json"
	"fmt"
)

type identityResponse struct {
	Accounts  []Account `json:"accounts"`
	Identity  `json:"identity"`
	Item      `json:"item"`
	RequestID string `json:"request_id"`
}

// Identity fetches ID info
func (i Item) Identity(clnt Client, accountIDs ...string) (Identity, error) {
	bts, err := get(fmt.Sprintf("%v/identity/get", clnt.envURL), clnt, i)
	if err != nil {
		return Identity{}, err
	}
	ir := identityResponse{}
	if err := json.Unmarshal(bts, &ir); err != nil {
		return Identity{}, err
	}
	return ir.Identity, nil
}
