package shared

type InvoiceObjectInvoiceTypeEnum string

const (
	InvoiceObjectInvoiceTypeEnumSubscription InvoiceObjectInvoiceTypeEnum = "subscription"
	InvoiceObjectInvoiceTypeEnumAddOn        InvoiceObjectInvoiceTypeEnum = "add_on"
	InvoiceObjectInvoiceTypeEnumCredit       InvoiceObjectInvoiceTypeEnum = "credit"
)

type InvoiceObjectPaymentStatusEnum string

const (
	InvoiceObjectPaymentStatusEnumPending   InvoiceObjectPaymentStatusEnum = "pending"
	InvoiceObjectPaymentStatusEnumSucceeded InvoiceObjectPaymentStatusEnum = "succeeded"
	InvoiceObjectPaymentStatusEnumFailed    InvoiceObjectPaymentStatusEnum = "failed"
)

type InvoiceObjectStatusEnum string

const (
	InvoiceObjectStatusEnumDraft     InvoiceObjectStatusEnum = "draft"
	InvoiceObjectStatusEnumFinalized InvoiceObjectStatusEnum = "finalized"
)

type InvoiceObject struct {
	AmountCents          *int64                          `json:"amount_cents,omitempty"`
	AmountCurrency       *string                         `json:"amount_currency,omitempty"`
	CreditAmountCents    *int64                          `json:"credit_amount_cents,omitempty"`
	CreditAmountCurrency *string                         `json:"credit_amount_currency,omitempty"`
	Credits              []CreditObject                  `json:"credits,omitempty"`
	Customer             *CustomerObject                 `json:"customer,omitempty"`
	Fees                 []FeeObject                     `json:"fees,omitempty"`
	FileURL              *string                         `json:"file_url,omitempty"`
	InvoiceType          *InvoiceObjectInvoiceTypeEnum   `json:"invoice_type,omitempty"`
	IssuingDate          *string                         `json:"issuing_date,omitempty"`
	LagoID               *string                         `json:"lago_id,omitempty"`
	Legacy               *bool                           `json:"legacy,omitempty"`
	Number               *string                         `json:"number,omitempty"`
	PaymentStatus        *InvoiceObjectPaymentStatusEnum `json:"payment_status,omitempty"`
	SequentialID         *int64                          `json:"sequential_id,omitempty"`
	Status               *InvoiceObjectStatusEnum        `json:"status,omitempty"`
	Subscriptions        []SubscriptionObject            `json:"subscriptions,omitempty"`
	TotalAmountCents     *int64                          `json:"total_amount_cents,omitempty"`
	TotalAmountCurrency  *string                         `json:"total_amount_currency,omitempty"`
	VatAmountCents       *int64                          `json:"vat_amount_cents,omitempty"`
	VatAmountCurrency    *string                         `json:"vat_amount_currency,omitempty"`
}
