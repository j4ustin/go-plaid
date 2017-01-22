package plaid

import "testing"

func TestGetBalance(t *testing.T) {
	bal := UseBalance(testConfiguration(t))
	accts, err := bal.GetBalance("test_wells")
	if err != nil {
		t.Fatal(err)
	}
	if len(accts) != 4 {
		t.Errorf("Expected to get 4 accounts back")
	}
}
