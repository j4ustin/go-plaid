package plaid

import (
	"encoding/json"
	"fmt"
)

type balanceResponse struct {
	Accounts  []Account `json:"accounts"`
	Item      `json:"item"`
	RequestID string `json:"request_id"`
}

// Balance is used to fetch balance info for a given item
func (i Item) Balance(clnt Client, accountIDs ...string) ([]Account, error) {
	bts, err := get(fmt.Sprintf("%v/balance/get", clnt.envURL), clnt, i)
	if err != nil {
		return nil, err
	}
	br := balanceResponse{}
	if err := json.Unmarshal(bts, &br); err != nil {
		return nil, err
	}
	return br.Accounts, nil
}
