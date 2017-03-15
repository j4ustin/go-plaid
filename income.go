package plaid

import (
	"encoding/json"
	"fmt"
)

type incomeResponse struct {
	Income    `json:"income"`
	Item      `json:"item"`
	RequestID string `json:"request_id"`
}

// Income fetches income info
func (i Item) income(clnt Client, accountIDs ...string) (Income, error) {
	bts, err := get(fmt.Sprintf("%v/income/get", clnt.envURL), clnt, i)
	if err != nil {
		return Income{}, err
	}
	ir := incomeResponse{}
	if err := json.Unmarshal(bts, &ir); err != nil {
		return Income{}, err
	}
	return ir.Income, nil
}
