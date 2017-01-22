package plaid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrMfaRequired is returned when MFA is required from the client
	ErrMfaRequired = errors.New("MFA Required")
)

// errorMessage is used to unmarshal error messages from Plaid
type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Resolve string `json:"resolve"`
}

// checkStatusCode handles status codes
func checkStatusCode(statusCode int, body []byte) error {
	switch statusCode {
	case http.StatusOK:
		return nil
	case http.StatusCreated:
		return ErrMfaRequired
	default:
		em := errorMessage{}
		if err := json.Unmarshal(body, &em); err != nil {
			return err
		}
		return fmt.Errorf("Message: " + em.Message + " Resolution: " + em.Resolve)
	}
}
