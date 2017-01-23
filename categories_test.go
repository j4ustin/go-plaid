package plaid

import "testing"

func TestFetchAll(t *testing.T) {
	cats, err := UseCategory(testConfiguration(t)).FetchAll()
	if err != nil {
		t.Fatal(err)
	}
	// TODO: Write a better test
	if len(cats) != 602 {
		t.Errorf("Did not get expected number of categories back")
	}
}

func TestFetchByID(t *testing.T) {
	cat, err := UseCategory(testConfiguration(t)).FetchByID("10000000")
	if err != nil {
		t.Fatal(err)
	}
	if cat.ID != "10000000" {
		t.Errorf("Expected to get back correct category")
	}
}
