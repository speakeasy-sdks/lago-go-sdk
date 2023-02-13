package shared

type EventObject struct {
	Code                   *string                `json:"code,omitempty"`
	CreatedAt              *string                `json:"created_at,omitempty"`
	ExternalCustomerID     *string                `json:"external_customer_id,omitempty"`
	ExternalSubscriptionID *string                `json:"external_subscription_id,omitempty"`
	LagoCustomerID         *string                `json:"lago_customer_id,omitempty"`
	LagoID                 *string                `json:"lago_id,omitempty"`
	LagoSubscriptionID     *string                `json:"lago_subscription_id,omitempty"`
	Properties             map[string]interface{} `json:"properties,omitempty"`
	Timestamp              *string                `json:"timestamp,omitempty"`
	TransactionID          *string                `json:"transaction_id,omitempty"`
}
