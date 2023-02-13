package shared

type WalletInputWallet struct {
	Currency           *string  `json:"currency,omitempty"`
	ExpirationAt       *string  `json:"expiration_at,omitempty"`
	ExternalCustomerID *string  `json:"external_customer_id,omitempty"`
	GrantedCredits     *float64 `json:"granted_credits,omitempty"`
	Name               *string  `json:"name,omitempty"`
	PaidCredits        *float64 `json:"paid_credits,omitempty"`
	RateAmount         *float64 `json:"rate_amount,omitempty"`
}

type WalletInput struct {
	Wallet *WalletInputWallet `json:"wallet,omitempty"`
}
