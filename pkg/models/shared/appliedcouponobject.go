package shared

type AppliedCouponObjectFrequencyEnum string

const (
	AppliedCouponObjectFrequencyEnumOnce      AppliedCouponObjectFrequencyEnum = "once"
	AppliedCouponObjectFrequencyEnumRecurring AppliedCouponObjectFrequencyEnum = "recurring"
)

type AppliedCouponObjectStatusEnum string

const (
	AppliedCouponObjectStatusEnumActive     AppliedCouponObjectStatusEnum = "active"
	AppliedCouponObjectStatusEnumTerminated AppliedCouponObjectStatusEnum = "terminated"
)

type AppliedCouponObject struct {
	AmountCents                *int64                            `json:"amount_cents,omitempty"`
	AmountCentsRemaining       *int64                            `json:"amount_cents_remaining,omitempty"`
	AmountCurrency             *string                           `json:"amount_currency,omitempty"`
	CouponCode                 *string                           `json:"coupon_code,omitempty"`
	CreatedAt                  *string                           `json:"created_at,omitempty"`
	ExpirationAt               *string                           `json:"expiration_at,omitempty"`
	ExternalCustomerID         *string                           `json:"external_customer_id,omitempty"`
	Frequency                  *AppliedCouponObjectFrequencyEnum `json:"frequency,omitempty"`
	FrequencyDuration          *int64                            `json:"frequency_duration,omitempty"`
	FrequencyDurationRemaining *int64                            `json:"frequency_duration_remaining,omitempty"`
	LagoCouponID               *string                           `json:"lago_coupon_id,omitempty"`
	LagoCustomerID             *string                           `json:"lago_customer_id,omitempty"`
	LagoID                     *string                           `json:"lago_id,omitempty"`
	PercentageRate             *float64                          `json:"percentage_rate,omitempty"`
	Status                     *AppliedCouponObjectStatusEnum    `json:"status,omitempty"`
	TerminatedAt               *string                           `json:"terminated_at,omitempty"`
}
