package shared

type CreditNoteItemObject struct {
	AmountCents    *int64     `json:"amount_cents,omitempty"`
	AmountCurrency *string    `json:"amount_currency,omitempty"`
	Fee            *FeeObject `json:"fee,omitempty"`
	LagoID         *string    `json:"lago_id,omitempty"`
}
