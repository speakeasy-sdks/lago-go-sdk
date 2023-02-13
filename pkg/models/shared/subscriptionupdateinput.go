package shared

type SubscriptionUpdateInputSubscription struct {
	Name           *string `json:"name,omitempty"`
	SubscriptionAt *string `json:"subscription_at,omitempty"`
}

type SubscriptionUpdateInput struct {
	Subscription *SubscriptionUpdateInputSubscription `json:"subscription,omitempty"`
}
