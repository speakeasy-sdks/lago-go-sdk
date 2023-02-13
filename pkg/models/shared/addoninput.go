package shared

type AddOnInputAddOn struct {
	AmountCents    *int64  `json:"amount_cents,omitempty"`
	AmountCurrency *string `json:"amount_currency,omitempty"`
	Code           *string `json:"code,omitempty"`
	Description    *string `json:"description,omitempty"`
	Name           *string `json:"name,omitempty"`
}

type AddOnInput struct {
	AddOn *AddOnInputAddOn `json:"add_on,omitempty"`
}
