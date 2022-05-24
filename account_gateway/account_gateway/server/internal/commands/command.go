package commands

type OpenAccountCommand struct {
	AccountType   int    `json:"account_type"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}
