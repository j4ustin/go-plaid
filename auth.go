package plaid

// Auth grants access to the Auth product
type Auth interface {
	AddMFA()
	GetData()
	AddUser()
	UpdateUser()
	DeleteUser()
}

// auth sets up an Auth service client
type auth struct {
	Client
	remote string
}

const (
	authURL = "/auth"
)

// UseAuth configures an auth product for use
func UseAuth(clnt Client) Auth {
	return &auth{
		remote: clnt.envURL + authURL,
		Client: clnt,
	}
}

func (a *auth) AddMFA()     {}
func (a *auth) GetData()    {}
func (a *auth) AddUser()    {}
func (a *auth) UpdateUser() {}
func (a *auth) DeleteUser() {}
