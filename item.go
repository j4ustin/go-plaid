package plaid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type itemResponse struct {
	AccessToken string `json:"access_token"`
	RequestID   string `json:"requestID"`
	Item        `json:"item"`
}

type rotateResponse struct {
	NewAccessToken string `json:"new_access_token"`
	RequestID      string `json:"request_id"`
}

type updateVersionRequest struct {
	ClientID      string `json:"client_id"`
	Secret        string `json:"secret"`
	AccessTokenV1 string `json:"access_token_v1"`
}

// GetItem fetches an item and is the first call made for the library
func GetItem(clnt Client, accessToken string) (Item, error) {
	bts, err := get(fmt.Sprintf("%v/item/get", clnt.envURL), clnt, Item{
		AccessToken: accessToken,
	})
	if err != nil {
		return Item{}, err
	}
	ir := itemResponse{}
	err = json.Unmarshal(bts, &ir)
	return ir.Item, err
}

// RotateAccessToken fetches an item and is the first call made for the library
func RotateAccessToken(clnt Client, accessToken string) (Item, error) {
	bts, err := get(fmt.Sprintf("%v/item/access_token/invalidate", clnt.envURL), clnt, Item{
		AccessToken: accessToken,
	})
	if err != nil {
		return Item{}, err
	}
	rr := rotateResponse{}
	if err = json.Unmarshal(bts, &rr); err != nil {
		return Item{}, err
	}
	return GetItem(clnt, rr.NewAccessToken)
}

// UpdateAccessTokenVersion updates legacy access tokens to use the new API
func UpdateAccessTokenVersion(clnt Client, accessToken string) (Item, error) {
	bts, err := json.Marshal(updateVersionRequest{
		ClientID:      clnt.clientID,
		Secret:        clnt.clientSecret,
		AccessTokenV1: accessToken,
	})
	if err != nil {
		return Item{}, err
	}
	res, err := http.Post(fmt.Sprintf("%v/item/access_token/update_version", clnt.envURL), "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return Item{}, err
	}
	bts, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return Item{}, err
	}
	if res.StatusCode != http.StatusOK {
		return Item{}, formatError(bts)
	}
	ir := itemResponse{}
	if err = json.Unmarshal(bts, &ir); err != nil {
		return Item{}, err
	}
	return GetItem(clnt, ir.AccessToken)
}

// Delete removes an Item
func (i Item) Delete(clnt Client) error {
	_, err := get(fmt.Sprintf("%v/item/delete", clnt.envURL), clnt, i)
	return err
}
