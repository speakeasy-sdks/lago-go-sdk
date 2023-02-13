package shared

type PlanInputPlanChargesChargeModelEnum string

const (
	PlanInputPlanChargesChargeModelEnumStandard   PlanInputPlanChargesChargeModelEnum = "standard"
	PlanInputPlanChargesChargeModelEnumGraduated  PlanInputPlanChargesChargeModelEnum = "graduated"
	PlanInputPlanChargesChargeModelEnumPackage    PlanInputPlanChargesChargeModelEnum = "package"
	PlanInputPlanChargesChargeModelEnumPercentage PlanInputPlanChargesChargeModelEnum = "percentage"
	PlanInputPlanChargesChargeModelEnumVolume     PlanInputPlanChargesChargeModelEnum = "volume"
)

type PlanInputPlanChargesGroupProperties struct {
	GroupID *string                `json:"group_id,omitempty"`
	Values  map[string]interface{} `json:"values,omitempty"`
}

type PlanInputPlanCharges struct {
	BillableMetricID *string                               `json:"billable_metric_id,omitempty"`
	ChargeModel      *PlanInputPlanChargesChargeModelEnum  `json:"charge_model,omitempty"`
	GroupProperties  []PlanInputPlanChargesGroupProperties `json:"group_properties,omitempty"`
	ID               *string                               `json:"id,omitempty"`
	Properties       map[string]interface{}                `json:"properties,omitempty"`
}

type PlanInputPlanIntervalEnum string

const (
	PlanInputPlanIntervalEnumWeekly  PlanInputPlanIntervalEnum = "weekly"
	PlanInputPlanIntervalEnumMonthly PlanInputPlanIntervalEnum = "monthly"
	PlanInputPlanIntervalEnumYearly  PlanInputPlanIntervalEnum = "yearly"
)

type PlanInputPlan struct {
	AmountCents        *int64                     `json:"amount_cents,omitempty"`
	AmountCurrency     *string                    `json:"amount_currency,omitempty"`
	BillChargesMonthly *bool                      `json:"bill_charges_monthly,omitempty"`
	Charges            []PlanInputPlanCharges     `json:"charges,omitempty"`
	Code               *string                    `json:"code,omitempty"`
	Description        *string                    `json:"description,omitempty"`
	Interval           *PlanInputPlanIntervalEnum `json:"interval,omitempty"`
	Name               *string                    `json:"name,omitempty"`
	PayInAdvance       *bool                      `json:"pay_in_advance,omitempty"`
	TrialPeriod        *float64                   `json:"trial_period,omitempty"`
}

type PlanInput struct {
	Plan *PlanInputPlan `json:"plan,omitempty"`
}
