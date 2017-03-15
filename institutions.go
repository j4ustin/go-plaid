package plaid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type institutionResponse struct {
	Institutions []Institution `json:"institutions"`
	RequestID    string        `json:"request_id"`
}

type idRequest struct {
	InstitutionID string `json:"institution_id"`
	PublicKey     string `json:"public_key"`
}

type idResponse struct {
	Institution `json:"institution"`
	RequestID   string `json:"request_id"`
}

type query struct {
	key       string
	parameter string
}

type offsetRequest struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
}

type searchRequest struct {
	Query     string   `json:"query"`
	Products  []string `json:"products"`
	PublicKey string   `json:"public_key"`
}

// Offsetter allows for a more dynamic query setter
type Offsetter func(*offsetRequest)

// GetInstitutions preforms a search on the institutions endpoint. The serach
// defaults to 500 institutions with no offset.
func GetInstitutions(clnt Client, ofst ...Offsetter) ([]Institution, error) {
	offset := offsetRequest{
		Count:  500,
		Offset: 1,
	}
	for _, ofs := range ofst {
		ofs(&offset)
	}
	bts, err := json.Marshal(offset)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(fmt.Sprintf("%v/institutions/get", clnt.envURL), "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	bts, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, formatError(bts)
	}
	ir := institutionResponse{}
	err = json.Unmarshal(bts, &ir)
	return ir.Institutions, err
}

// SetOffset is used to overwrite the default options for the GetInstitutions method
func SetOffset(count, offset int) Offsetter {
	return func(ofs *offsetRequest) {
		ofs.Count = count
		ofs.Offset = offset
	}
}

// InstitutionSearch performs a search on institutions
func InstitutionSearch(clnt Client, publicKey, query string, products ...string) ([]Institution, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("Require at least one product to search for institutions")
	}
	var prds []string
	for _, prd := range products {
		prds = append(prds, prd)
	}
	bts, err := json.Marshal(searchRequest{
		Query:     query,
		PublicKey: publicKey,
		Products:  prds,
	})
	if err != nil {
		return nil, err
	}
	res, err := http.Post(fmt.Sprintf("%v/institutions/search", clnt.envURL), "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	bts, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, formatError(bts)
	}
	ir := institutionResponse{}
	err = json.Unmarshal(bts, &ir)
	return ir.Institutions, err
}

// InstitutionByID finds a single institution by its ID
func InstitutionByID(clnt Client, publicKey, id string) (Institution, error) {
	if id == "" {
		return Institution{}, fmt.Errorf("Requre an ID in order to perform search")
	}
	bts, err := json.Marshal(idRequest{})
	if err != nil {
		return Institution{}, err
	}
	res, err := http.Post(fmt.Sprintf("%v/institutions/get_by_id", clnt.envURL), "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return Institution{}, err
	}
	bts, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return Institution{}, err
	}
	if res.StatusCode != http.StatusOK {
		return Institution{}, formatError(bts)
	}
	ir := idResponse{}
	err = json.Unmarshal(bts, &ir)
	return ir.Institution, err
}
