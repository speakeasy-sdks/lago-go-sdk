package shared

type AddOnObject struct {
	AmountCents    *int64  `json:"amount_cents,omitempty"`
	AmountCurrency *string `json:"amount_currency,omitempty"`
	Code           *string `json:"code,omitempty"`
	CreatedAt      *string `json:"created_at,omitempty"`
	Description    *string `json:"description,omitempty"`
	LagoID         *string `json:"lago_id,omitempty"`
	Name           *string `json:"name,omitempty"`
}
