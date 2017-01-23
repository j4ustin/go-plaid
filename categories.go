package plaid

import (
	"encoding/json"
	"fmt"
)

// Categories grants access to the categories product
type Categories interface {
	FetchAll() ([]Category, error)
	FetchByID(string) (Category, error)
}

// Category stores information about a category
type Category struct {
	ID        string
	Type      string
	Hierarchy []string
}

type categories struct {
	Client
	remote string
}

const (
	catURL = "/categories"
)

// UseCategory configures an auth product for use
func UseCategory(clnt Client) Categories {
	return &categories{
		remote: clnt.envURL + catURL,
		Client: clnt,
	}
}

// FetchAll returns all categories
func (c *categories) FetchAll() ([]Category, error) {
	res, err := get(c.remote)
	if err != nil {
		return nil, err
	}
	var cts []Category
	return cts, json.Unmarshal(res, &cts)
}

// FetchByID returns a single cateogry
func (c *categories) FetchByID(id string) (Category, error) {
	res, err := get(fmt.Sprintf("%v/%v", c.remote, id))
	if err != nil {
		return Category{}, err
	}
	ctg := Category{}
	return ctg, json.Unmarshal(res, &ctg)
}
