package shared

type CreditNoteInputCreditNoteItems struct {
	AmountCents *int64  `json:"amount_cents,omitempty"`
	FeeID       *string `json:"fee_id,omitempty"`
}

type CreditNoteInputCreditNoteReasonEnum string

const (
	CreditNoteInputCreditNoteReasonEnumDuplicatedCharge      CreditNoteInputCreditNoteReasonEnum = "duplicated_charge"
	CreditNoteInputCreditNoteReasonEnumProductUnsatisfactory CreditNoteInputCreditNoteReasonEnum = "product_unsatisfactory"
	CreditNoteInputCreditNoteReasonEnumOrderChange           CreditNoteInputCreditNoteReasonEnum = "order_change"
	CreditNoteInputCreditNoteReasonEnumOrderCancellation     CreditNoteInputCreditNoteReasonEnum = "order_cancellation"
	CreditNoteInputCreditNoteReasonEnumFraudulentCharge      CreditNoteInputCreditNoteReasonEnum = "fraudulent_charge"
	CreditNoteInputCreditNoteReasonEnumOther                 CreditNoteInputCreditNoteReasonEnum = "other"
)

type CreditNoteInputCreditNote struct {
	CreditAmountCents *int64                               `json:"credit_amount_cents,omitempty"`
	Description       *string                              `json:"description,omitempty"`
	InvoiceID         *string                              `json:"invoice_id,omitempty"`
	Items             []CreditNoteInputCreditNoteItems     `json:"items,omitempty"`
	Reason            *CreditNoteInputCreditNoteReasonEnum `json:"reason,omitempty"`
	RefundAmountCents *int64                               `json:"refund_amount_cents,omitempty"`
}

type CreditNoteInput struct {
	CreditNote *CreditNoteInputCreditNote `json:"credit_note,omitempty"`
}
