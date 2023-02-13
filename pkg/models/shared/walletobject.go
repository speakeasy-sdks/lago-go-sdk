package shared

type WalletObjectStatusEnum string

const (
	WalletObjectStatusEnumActive     WalletObjectStatusEnum = "active"
	WalletObjectStatusEnumTerminated WalletObjectStatusEnum = "terminated"
)

type WalletObject struct {
	Balance              *float64                `json:"balance,omitempty"`
	ConsumedCredits      *float64                `json:"consumed_credits,omitempty"`
	CreatedAt            *string                 `json:"created_at,omitempty"`
	CreditsBalance       *float64                `json:"credits_balance,omitempty"`
	Currency             *string                 `json:"currency,omitempty"`
	ExpirationAt         *string                 `json:"expiration_at,omitempty"`
	ExternalCustomerID   *string                 `json:"external_customer_id,omitempty"`
	LagoCustomerID       *string                 `json:"lago_customer_id,omitempty"`
	LagoID               *string                 `json:"lago_id,omitempty"`
	LastBalanceSyncAt    *string                 `json:"last_balance_sync_at,omitempty"`
	LastConsumedCreditAt *string                 `json:"last_consumed_credit_at,omitempty"`
	Name                 *string                 `json:"name,omitempty"`
	RateAmount           *float64                `json:"rate_amount,omitempty"`
	Status               *WalletObjectStatusEnum `json:"status,omitempty"`
	TerminatedAt         *string                 `json:"terminated_at,omitempty"`
}
