package shared

type CreditNoteObjectCreditStatusEnum string

const (
	CreditNoteObjectCreditStatusEnumAvailable CreditNoteObjectCreditStatusEnum = "available"
	CreditNoteObjectCreditStatusEnumConsumed  CreditNoteObjectCreditStatusEnum = "consumed"
	CreditNoteObjectCreditStatusEnumVoided    CreditNoteObjectCreditStatusEnum = "voided"
)

type CreditNoteObjectReasonEnum string

const (
	CreditNoteObjectReasonEnumDuplicatedCharge      CreditNoteObjectReasonEnum = "duplicated_charge"
	CreditNoteObjectReasonEnumProductUnsatisfactory CreditNoteObjectReasonEnum = "product_unsatisfactory"
	CreditNoteObjectReasonEnumOrderChange           CreditNoteObjectReasonEnum = "order_change"
	CreditNoteObjectReasonEnumOrderCancellation     CreditNoteObjectReasonEnum = "order_cancellation"
	CreditNoteObjectReasonEnumFraudulentCharge      CreditNoteObjectReasonEnum = "fraudulent_charge"
	CreditNoteObjectReasonEnumOther                 CreditNoteObjectReasonEnum = "other"
)

type CreditNoteObjectRefundStatusEnum string

const (
	CreditNoteObjectRefundStatusEnumPending   CreditNoteObjectRefundStatusEnum = "pending"
	CreditNoteObjectRefundStatusEnumSucceeded CreditNoteObjectRefundStatusEnum = "succeeded"
	CreditNoteObjectRefundStatusEnumFailed    CreditNoteObjectRefundStatusEnum = "failed"
)

type CreditNoteObject struct {
	BalanceAmountCents                *int64                            `json:"balance_amount_cents,omitempty"`
	BalanceAmountCurrency             *string                           `json:"balance_amount_currency,omitempty"`
	CreatedAt                         *string                           `json:"created_at,omitempty"`
	CreditAmountCents                 *int64                            `json:"credit_amount_cents,omitempty"`
	CreditAmountCurrency              *string                           `json:"credit_amount_currency,omitempty"`
	CreditStatus                      *CreditNoteObjectCreditStatusEnum `json:"credit_status,omitempty"`
	Customer                          *CustomerObject                   `json:"customer,omitempty"`
	Description                       *string                           `json:"description,omitempty"`
	FileURL                           *string                           `json:"file_url,omitempty"`
	InvoiceNumber                     *string                           `json:"invoice_number,omitempty"`
	IssuingDate                       *string                           `json:"issuing_date,omitempty"`
	Items                             []CreditNoteItemObject            `json:"items,omitempty"`
	LagoID                            *string                           `json:"lago_id,omitempty"`
	LagoInvoiceID                     *string                           `json:"lago_invoice_id,omitempty"`
	Number                            *string                           `json:"number,omitempty"`
	Reason                            *CreditNoteObjectReasonEnum       `json:"reason,omitempty"`
	RefundAmountCents                 *int64                            `json:"refund_amount_cents,omitempty"`
	RefundAmountCurrency              *string                           `json:"refund_amount_currency,omitempty"`
	RefundStatus                      *CreditNoteObjectRefundStatusEnum `json:"refund_status,omitempty"`
	SequentialID                      *int64                            `json:"sequential_id,omitempty"`
	SubTotalVatExcludedAmountCents    *int64                            `json:"sub_total_vat_excluded_amount_cents,omitempty"`
	SubTotalVatExcludedAmountCurrency *string                           `json:"sub_total_vat_excluded_amount_currency,omitempty"`
	TotalAmountCents                  *int64                            `json:"total_amount_cents,omitempty"`
	TotalAmountCurrency               *string                           `json:"total_amount_currency,omitempty"`
	UpdatedAt                         *string                           `json:"updated_at,omitempty"`
	VatAmountCents                    *int64                            `json:"vat_amount_cents,omitempty"`
	VatAmountCurrency                 *string                           `json:"vat_amount_currency,omitempty"`
}
