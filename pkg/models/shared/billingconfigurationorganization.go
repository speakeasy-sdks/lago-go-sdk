package shared

type BillingConfigurationOrganization struct {
	InvoiceFooter      *string  `json:"invoice_footer,omitempty"`
	InvoiceGracePeriod *int64   `json:"invoice_grace_period,omitempty"`
	VatRate            *float64 `json:"vat_rate,omitempty"`
}
