package plaid

import (
	"bytes"
	"encoding/json"
)

// Balance grants access to the Balance product
type Balance interface {
	GetBalance(string) ([]Account, error)
}

type balance struct {
	Client
	remote string
}

const (
	balanceURL = "/balance"
)

// UseBalance configures an balance product for use
func UseBalance(clnt Client) Balance {
	return &balance{
		remote: clnt.envURL + balanceURL,
		Client: clnt,
	}
}

func (b *balance) GetBalance(accessToken string) ([]Account, error) {
	bts, err := json.Marshal(getRequest{
		AccessToken: accessToken,
		ClientID:    b.clientID,
		Secret:      b.clientSecret,
	})
	if err != nil {
		return nil, err
	}
	res, err := post(b.remote, bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	ar := accountsRes{}
	if err := json.Unmarshal(res, &ar); err != nil {
		return nil, err
	}
	return ar.Accounts, nil
}
