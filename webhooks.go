package plaid

// Receiver is a function that is passed into the webhook
type Receiver func()

// WebhookResponse is returnd on a webhook route. It is parsed and the
type WebhookResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
	Code        int    `json:"code"`
}
