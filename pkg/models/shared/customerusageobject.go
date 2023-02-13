package shared

type CustomerUsageObject struct {
	AmountCents         *int64              `json:"amount_cents,omitempty"`
	AmountCurrency      *string             `json:"amount_currency,omitempty"`
	ChargesUsage        []ChargeUsageObject `json:"charges_usage,omitempty"`
	FromDatetime        *string             `json:"from_datetime,omitempty"`
	IssuingDate         *string             `json:"issuing_date,omitempty"`
	ToDatetime          *string             `json:"to_datetime,omitempty"`
	TotalAmountCents    *int64              `json:"total_amount_cents,omitempty"`
	TotalAmountCurrency *string             `json:"total_amount_currency,omitempty"`
	VatAmountCents      *int64              `json:"vat_amount_cents,omitempty"`
	VatAmountCurrency   *string             `json:"vat_amount_currency,omitempty"`
}
