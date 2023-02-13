package shared

type BatchEventInputEvent struct {
	Code                    *string                `json:"code,omitempty"`
	ExternalCustomerID      *string                `json:"external_customer_id,omitempty"`
	ExternalSubscriptionIds []string               `json:"external_subscription_ids,omitempty"`
	Properties              map[string]interface{} `json:"properties,omitempty"`
	Timestamp               *int64                 `json:"timestamp,omitempty"`
	TransactionID           *string                `json:"transaction_id,omitempty"`
}

type BatchEventInput struct {
	Event *BatchEventInputEvent `json:"event,omitempty"`
}
