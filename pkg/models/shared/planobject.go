package shared

type PlanObjectIntervalEnum string

const (
	PlanObjectIntervalEnumWeekly  PlanObjectIntervalEnum = "weekly"
	PlanObjectIntervalEnumMonthly PlanObjectIntervalEnum = "monthly"
	PlanObjectIntervalEnumYearly  PlanObjectIntervalEnum = "yearly"
)

type PlanObject struct {
	AmountCents        *int64                  `json:"amount_cents,omitempty"`
	AmountCurrency     *string                 `json:"amount_currency,omitempty"`
	BillChargesMonthly *bool                   `json:"bill_charges_monthly,omitempty"`
	Charges            []ChargeObject          `json:"charges,omitempty"`
	Code               *string                 `json:"code,omitempty"`
	CreatedAt          *string                 `json:"created_at,omitempty"`
	Description        *string                 `json:"description,omitempty"`
	Interval           *PlanObjectIntervalEnum `json:"interval,omitempty"`
	LagoID             *string                 `json:"lago_id,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	PayInAdvance       *bool                   `json:"pay_in_advance,omitempty"`
	TrialPeriod        *float64                `json:"trial_period,omitempty"`
}
