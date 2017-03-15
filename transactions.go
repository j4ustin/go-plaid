package plaid

import (
	"encoding/json"
	"fmt"
)

type transactionResponse struct {
	Item         `json:"item"`
	RequestID    string        `json:"request_id"`
	Accounts     []Account     `json:"accounts"`
	Transactions []Transaction `json:"transactions"`
}

// Transactions returns transactions associated with the access token
func (i Item) Transactions(clnt Client) ([]Transaction, error) {
	bts, err := get(fmt.Sprintf("%v/transactions/get", clnt.envURL), clnt, i)
	if err != nil {
		return nil, err
	}
	tr := transactionResponse{}
	err = json.Unmarshal(bts, &tr)
	return tr.Transactions, err
}
