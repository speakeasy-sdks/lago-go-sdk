package shared

type CreditObjectItem struct {
	Code   *string `json:"code,omitempty"`
	LagoID *string `json:"lago_id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Type   *string `json:"type,omitempty"`
}

type CreditObject struct {
	AmountCents    *int64            `json:"amount_cents,omitempty"`
	AmountCurrency *string           `json:"amount_currency,omitempty"`
	Item           *CreditObjectItem `json:"item,omitempty"`
	LagoID         *string           `json:"lago_id,omitempty"`
}
