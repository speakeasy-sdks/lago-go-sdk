package shared

type AppliedCouponInputAppliedCouponFrequencyEnum string

const (
	AppliedCouponInputAppliedCouponFrequencyEnumOnce      AppliedCouponInputAppliedCouponFrequencyEnum = "once"
	AppliedCouponInputAppliedCouponFrequencyEnumRecurring AppliedCouponInputAppliedCouponFrequencyEnum = "recurring"
)

type AppliedCouponInputAppliedCoupon struct {
	AmountCents        *int64                                        `json:"amount_cents,omitempty"`
	AmountCurrency     *string                                       `json:"amount_currency,omitempty"`
	CouponCode         *string                                       `json:"coupon_code,omitempty"`
	ExternalCustomerID *string                                       `json:"external_customer_id,omitempty"`
	Frequency          *AppliedCouponInputAppliedCouponFrequencyEnum `json:"frequency,omitempty"`
	FrequencyDuration  *int64                                        `json:"frequency_duration,omitempty"`
	PercentageRate     *float64                                      `json:"percentage_rate,omitempty"`
}

type AppliedCouponInput struct {
	AppliedCoupon *AppliedCouponInputAppliedCoupon `json:"applied_coupon,omitempty"`
}
