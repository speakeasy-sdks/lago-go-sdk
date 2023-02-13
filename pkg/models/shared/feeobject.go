package shared

type FeeObjectItemTypeEnum string

const (
	FeeObjectItemTypeEnumCharge       FeeObjectItemTypeEnum = "charge"
	FeeObjectItemTypeEnumAddOn        FeeObjectItemTypeEnum = "add_on"
	FeeObjectItemTypeEnumSubscription FeeObjectItemTypeEnum = "subscription"
	FeeObjectItemTypeEnumCredit       FeeObjectItemTypeEnum = "credit"
)

type FeeObjectItem struct {
	Code *string                `json:"code,omitempty"`
	Name *string                `json:"name,omitempty"`
	Type *FeeObjectItemTypeEnum `json:"type,omitempty"`
}

type FeeObject struct {
	AmountCents       *int64         `json:"amount_cents,omitempty"`
	AmountCurrency    *string        `json:"amount_currency,omitempty"`
	EventsCount       *int64         `json:"events_count,omitempty"`
	Item              *FeeObjectItem `json:"item,omitempty"`
	LagoGroupID       *string        `json:"lago_group_id,omitempty"`
	LagoID            *string        `json:"lago_id,omitempty"`
	Units             *float64       `json:"units,omitempty"`
	VatAmountCents    *int64         `json:"vat_amount_cents,omitempty"`
	VatAmountCurrency *string        `json:"vat_amount_currency,omitempty"`
}
