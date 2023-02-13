package shared

type InvoiceInputInvoicePaymentStatusEnum string

const (
	InvoiceInputInvoicePaymentStatusEnumPending   InvoiceInputInvoicePaymentStatusEnum = "pending"
	InvoiceInputInvoicePaymentStatusEnumSucceeded InvoiceInputInvoicePaymentStatusEnum = "succeeded"
	InvoiceInputInvoicePaymentStatusEnumFailed    InvoiceInputInvoicePaymentStatusEnum = "failed"
)

type InvoiceInputInvoice struct {
	PaymentStatus *InvoiceInputInvoicePaymentStatusEnum `json:"payment_status,omitempty"`
}

type InvoiceInput struct {
	Invoice *InvoiceInputInvoice `json:"invoice,omitempty"`
}
