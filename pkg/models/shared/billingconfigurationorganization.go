package shared

type BillingConfigurationOrganization struct {
	DocumentLocale     *string  `json:"document_locale,omitempty"`
	InvoiceFooter      *string  `json:"invoice_footer,omitempty"`
	InvoiceGracePeriod *int64   `json:"invoice_grace_period,omitempty"`
	VatRate            *float64 `json:"vat_rate,omitempty"`
}
