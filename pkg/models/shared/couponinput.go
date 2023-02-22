package shared

type CouponInputCouponAppliesTo struct {
	PlanCodes []string `json:"plan_codes,omitempty"`
}

type CouponInputCouponCouponTypeEnum string

const (
	CouponInputCouponCouponTypeEnumFixedAmount CouponInputCouponCouponTypeEnum = "fixed_amount"
	CouponInputCouponCouponTypeEnumPercentage  CouponInputCouponCouponTypeEnum = "percentage"
)

type CouponInputCouponExpirationEnum string

const (
	CouponInputCouponExpirationEnumNoExpiration CouponInputCouponExpirationEnum = "no_expiration"
	CouponInputCouponExpirationEnumTimeLimit    CouponInputCouponExpirationEnum = "time_limit"
)

type CouponInputCouponFrequencyEnum string

const (
	CouponInputCouponFrequencyEnumOnce      CouponInputCouponFrequencyEnum = "once"
	CouponInputCouponFrequencyEnumRecurring CouponInputCouponFrequencyEnum = "recurring"
)

type CouponInputCoupon struct {
	AmountCents       *int64                           `json:"amount_cents,omitempty"`
	AmountCurrency    *string                          `json:"amount_currency,omitempty"`
	AppliesTo         *CouponInputCouponAppliesTo      `json:"applies_to,omitempty"`
	Code              *string                          `json:"code,omitempty"`
	CouponType        *CouponInputCouponCouponTypeEnum `json:"coupon_type,omitempty"`
	Expiration        *CouponInputCouponExpirationEnum `json:"expiration,omitempty"`
	ExpirationAt      *string                          `json:"expiration_at,omitempty"`
	Frequency         *CouponInputCouponFrequencyEnum  `json:"frequency,omitempty"`
	FrequencyDuration *int64                           `json:"frequency_duration,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	PercentageRate    *float64                         `json:"percentage_rate,omitempty"`
	Reusable          *bool                            `json:"reusable,omitempty"`
}

type CouponInput struct {
	Coupon *CouponInputCoupon `json:"coupon,omitempty"`
}
