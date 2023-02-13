package shared

type SubscriptionObjectBillingTimeEnum string

const (
	SubscriptionObjectBillingTimeEnumCalendar    SubscriptionObjectBillingTimeEnum = "calendar"
	SubscriptionObjectBillingTimeEnumAnniversary SubscriptionObjectBillingTimeEnum = "anniversary"
)

type SubscriptionObjectStatusEnum string

const (
	SubscriptionObjectStatusEnumActive     SubscriptionObjectStatusEnum = "active"
	SubscriptionObjectStatusEnumPending    SubscriptionObjectStatusEnum = "pending"
	SubscriptionObjectStatusEnumTerminated SubscriptionObjectStatusEnum = "terminated"
	SubscriptionObjectStatusEnumCanceled   SubscriptionObjectStatusEnum = "canceled"
)

type SubscriptionObject struct {
	BillingTime        *SubscriptionObjectBillingTimeEnum `json:"billing_time,omitempty"`
	CanceledAt         *string                            `json:"canceled_at,omitempty"`
	CreatedAt          *string                            `json:"created_at,omitempty"`
	DowngradePlanDate  *string                            `json:"downgrade_plan_date,omitempty"`
	ExternalCustomerID *string                            `json:"external_customer_id,omitempty"`
	ExternalID         *string                            `json:"external_id,omitempty"`
	LagoCustomerID     *string                            `json:"lago_customer_id,omitempty"`
	LagoID             *string                            `json:"lago_id,omitempty"`
	Name               *string                            `json:"name,omitempty"`
	NextPlanCode       *string                            `json:"next_plan_code,omitempty"`
	PlanCode           *string                            `json:"plan_code,omitempty"`
	PreviousPlanCode   *string                            `json:"previous_plan_code,omitempty"`
	StartedAt          *string                            `json:"started_at,omitempty"`
	Status             *SubscriptionObjectStatusEnum      `json:"status,omitempty"`
	SubscriptionAt     *string                            `json:"subscription_at,omitempty"`
	TerminatedAt       *string                            `json:"terminated_at,omitempty"`
}
