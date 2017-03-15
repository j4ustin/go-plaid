package plaid

import "encoding/json"

// Error is a formatted error type for plaid
type Error struct {
	ErrorType      string `json:"error_type"`
	HTTPCode       int    `json:"http_code"`
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	DisplayMessage string `json:"display_message"`
	RequestID      string `json:"request_id"`
}

// Error returns the user friendly display message and satisfies the
// error type
func (e Error) Error() string {
	return e.DisplayMessage
}

// formatError unmarshals a response into the plaid error type
func formatError(errMessage []byte) Error {
	pldErr := Error{}
	if err := json.Unmarshal(errMessage, &pldErr); err != nil {
		pldErr.DisplayMessage = err.Error()
	}
	return pldErr
}
