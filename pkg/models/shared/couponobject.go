package shared

type CouponObjectCouponTypeEnum string

const (
	CouponObjectCouponTypeEnumFixedAmount CouponObjectCouponTypeEnum = "fixed_amount"
	CouponObjectCouponTypeEnumPercentage  CouponObjectCouponTypeEnum = "percentage"
)

type CouponObjectExpirationEnum string

const (
	CouponObjectExpirationEnumNoExpiration CouponObjectExpirationEnum = "no_expiration"
	CouponObjectExpirationEnumTimeLimit    CouponObjectExpirationEnum = "time_limit"
)

type CouponObjectFrequencyEnum string

const (
	CouponObjectFrequencyEnumOnce      CouponObjectFrequencyEnum = "once"
	CouponObjectFrequencyEnumRecurring CouponObjectFrequencyEnum = "recurring"
)

type CouponObject struct {
	AmountCents       *int64                      `json:"amount_cents,omitempty"`
	AmountCurrency    *string                     `json:"amount_currency,omitempty"`
	Code              *string                     `json:"code,omitempty"`
	CouponType        *CouponObjectCouponTypeEnum `json:"coupon_type,omitempty"`
	CreatedAt         *string                     `json:"created_at,omitempty"`
	Expiration        *CouponObjectExpirationEnum `json:"expiration,omitempty"`
	ExpirationAt      *string                     `json:"expiration_at,omitempty"`
	Frequency         *CouponObjectFrequencyEnum  `json:"frequency,omitempty"`
	FrequencyDuration *int64                      `json:"frequency_duration,omitempty"`
	LagoID            *string                     `json:"lago_id,omitempty"`
	Name              *string                     `json:"name,omitempty"`
	PercentageRate    *float64                    `json:"percentage_rate,omitempty"`
	Reusable          *bool                       `json:"reusable,omitempty"`
}
