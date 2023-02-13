package shared

type EventInputEvent struct {
	Code                   *string                `json:"code,omitempty"`
	ExternalCustomerID     *string                `json:"external_customer_id,omitempty"`
	ExternalSubscriptionID *string                `json:"external_subscription_id,omitempty"`
	Properties             map[string]interface{} `json:"properties,omitempty"`
	Timestamp              *int64                 `json:"timestamp,omitempty"`
	TransactionID          *string                `json:"transaction_id,omitempty"`
}

type EventInput struct {
	Event *EventInputEvent `json:"event,omitempty"`
}
