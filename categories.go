package plaid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type categoryResponse struct {
	Categories []Category `json:"categories"`
	RequestID  string     `json:"requestID"`
}

// GetAllCategories fetches all categories
func GetAllCategories(clnt Client) ([]Category, error) {
	res, err := http.Post(fmt.Sprintf("%v/categories/get", clnt.envURL), "application/json", nil)
	if err != nil {
		return nil, err
	}
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	cr := categoryResponse{}
	err = json.Unmarshal(bts, &cr)
	return cr.Categories, err
}
