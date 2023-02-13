package shared

type CreditNoteUpdateInputCreditNoteRefundStatusEnum string

const (
	CreditNoteUpdateInputCreditNoteRefundStatusEnumPending   CreditNoteUpdateInputCreditNoteRefundStatusEnum = "pending"
	CreditNoteUpdateInputCreditNoteRefundStatusEnumSucceeded CreditNoteUpdateInputCreditNoteRefundStatusEnum = "succeeded"
	CreditNoteUpdateInputCreditNoteRefundStatusEnumFailed    CreditNoteUpdateInputCreditNoteRefundStatusEnum = "failed"
)

type CreditNoteUpdateInputCreditNote struct {
	RefundStatus *CreditNoteUpdateInputCreditNoteRefundStatusEnum `json:"refund_status,omitempty"`
}

type CreditNoteUpdateInput struct {
	CreditNote *CreditNoteUpdateInputCreditNote `json:"credit_note,omitempty"`
}
