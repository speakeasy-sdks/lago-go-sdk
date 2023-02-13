package shared

type SubscriptionCreateInputSubscriptionBillingTimeEnum string

const (
	SubscriptionCreateInputSubscriptionBillingTimeEnumCalendar    SubscriptionCreateInputSubscriptionBillingTimeEnum = "calendar"
	SubscriptionCreateInputSubscriptionBillingTimeEnumAnniversary SubscriptionCreateInputSubscriptionBillingTimeEnum = "anniversary"
)

type SubscriptionCreateInputSubscription struct {
	BillingTime        *SubscriptionCreateInputSubscriptionBillingTimeEnum `json:"billing_time,omitempty"`
	ExternalCustomerID *string                                             `json:"external_customer_id,omitempty"`
	ExternalID         *string                                             `json:"external_id,omitempty"`
	Name               *string                                             `json:"name,omitempty"`
	PlanCode           *string                                             `json:"plan_code,omitempty"`
	SubscriptionAt     *string                                             `json:"subscription_at,omitempty"`
}

type SubscriptionCreateInput struct {
	Subscription *SubscriptionCreateInputSubscription `json:"subscription,omitempty"`
}
