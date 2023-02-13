package shared

type AppliedAddOnObject struct {
	AddOnCode          *string `json:"add_on_code,omitempty"`
	AmountCents        *int64  `json:"amount_cents,omitempty"`
	AmountCurrency     *string `json:"amount_currency,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	ExternalCustomerID *string `json:"external_customer_id,omitempty"`
	LagoAddOnID        *string `json:"lago_add_on_id,omitempty"`
	LagoCustomerID     *string `json:"lago_customer_id,omitempty"`
	LagoID             *string `json:"lago_id,omitempty"`
}
