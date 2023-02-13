package shared

type AppliedAddOnInputAppliedAddOn struct {
	AddOnCode          *string `json:"add_on_code,omitempty"`
	AmountCents        *int64  `json:"amount_cents,omitempty"`
	AmountCurrency     *string `json:"amount_currency,omitempty"`
	ExternalCustomerID *string `json:"external_customer_id,omitempty"`
}

type AppliedAddOnInput struct {
	AppliedAddOn *AppliedAddOnInputAppliedAddOn `json:"applied_add_on,omitempty"`
}
