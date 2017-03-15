package plaid

// Item is the core type used to make all other requests
type Item struct {
	AccessToken       string
	AvailableProducts []string `json:"available_products"`
	BilledProducts    []string `json:"billed_products"`
	Err               Error    `json:"error"`
	InstitutionID     string   `json:"institution_id"`
	ItemID            string   `json:"item_id"`
	Webhook           string   `json:"webhook"`
}

// Category stores information about a category
type Category struct {
	ID        string
	Type      string
	Hierarchy []string
}

// Auth contains both accounts and numbers and is returnd from the auth call
type Auth struct {
	Accounts []Account `json:"accounts"`
	Numbers  []Number  `json:"numbers"`
}

// Account contains specific account info
type Account struct {
	AccountID string `json:"account_id"`
	Balances  struct {
		Available float64 `json:"available"`
		Current   float64 `json:"current"`
		Limit     float64 `json:"limit"`
	} `json:"balances"`
	Mask         string `json:"mask"`
	Name         string `json:"name"`
	OfficialName string `json:"official_name"`
	SubType      string `json:"subtype"`
	Type         string `json:"type"`
}

// Identity contains ID info
type Identity struct {
	Addresses []struct {
		Accounts []string `json:"accounts"`
		Data     struct {
			City   string `json:"city"`
			State  string `json:"state"`
			Street string `json:"street"`
			Zip    string `json:"zip"`
		} `json:"data"`
		Primary bool `json:"primary"`
	} `json:"addresses"`
	Emails []struct {
		Data    string `json:"data"`
		Primary bool   `json:"primary"`
		Type    string `json:"type"`
	}
	Names        []string `json:"names"`
	PhoneNumbers []struct {
		Primary bool   `json:"primary"`
		Type    string `json:"type"`
		Data    string `json:"data"`
	}
	Item      string `json:"item"`
	RequestID string `json:"request_id"`
}

// Transaction represents a single transaction
type Transaction struct {
	AccountID  string   `json:"account_id"`
	Amount     float64  `json:"amount"`
	Category   []string `json:"category"`
	CategoryID string   `json:"category_id"`
	Date       string   `json:"date"`
	Location   struct {
		Address     string `json:"address"`
		City        string `json:"city"`
		State       string `json:"state"`
		Zip         string `json:"zip"`
		Coordinates struct {
			Lat string `json:"lat"`
			Lon string `json:"lon"`
		} `json:"coordinates"`
		Name string `json:"name"`
		// PaymentMeta TODO: figure out object
		Pending              bool   `json:"pending"`
		PendingTransactionID string `json:"pending_transaction_id"`
		TransactionID        string `json:"transaction_id"`
		TransactionType      string `json:"transaction_type"`
	} `json:"location"`
}

// Institution represents one institution objects
type Institution struct {
	Credentials struct {
		Label string `json:"label"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"credentials"`
	HasMFA        bool     `json:"has_mfa"`
	InstitutionID string   `json:"institution_id"`
	MFA           []string `json:"mfa"`
	Name          string   `json:"name"`
	Products      []string `json:"products"`
}

// Income returns the income type for an item
type Income struct {
	Item   string `json:"item"`
	Income struct {
		IncomeStreams []struct {
			Confidence    int    `json:"confidence"`
			Days          int    `json:"days"`
			MonthlyIncome int    `json:"monthly_income"`
			Name          string `json:"name"`
		} `json:"income_streams"`
		LastYearIncome                      int `json:"last_year_income"`
		LastYearIncomeBeforeTax             int `json:"last_year_income_before_tax"`
		ProjectedYearlyIncome               int `json:"projected_yearly_income"`
		ProjectedYearlyIncomeBeforeTax      int `json:"projected_yearly_income_before_tax"`
		MaxNumberOfOverlappingIncomeStreams int `json:"max_number_of_overlapping_income_streams"`
		NumberOfIncomeStreams               int `json:"number_of_income_streams"`
	} `json:"income"`
	RequestID string `json:"request_id"`
}

// Number contains account numbers
type Number struct {
	Account     string `json:"account"`
	AccountID   string `json:"accountID"`
	Routing     string `json:"routing"`
	WireRouting string `json:"wire_routing"`
}
