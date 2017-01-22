package plaid

import (
	"encoding/json"
	"testing"
)

func TestCheckStatusCode(t *testing.T) {
	if err := checkStatusCode(200, nil); err != nil {
		t.Fatal(err)
	}
	if err := checkStatusCode(201, nil); err != ErrMfaRequired {
		t.Fatal(err)
	}
	bts, err := json.Marshal(errorMessage{
		Code:    1000,
		Message: "bad things happened",
		Resolve: "resolve accordingly",
	})
	err = checkStatusCode(401, bts)
	if err == nil || err.Error() != "Message: bad things happened Resolution: resolve accordingly" {
		t.Errorf("Expected correct  error message to be constructed")
	}
}
