package shared

type CustomerObject struct {
	AddressLine1         *string                `json:"address_line1,omitempty"`
	AddressLine2         *string                `json:"address_line2,omitempty"`
	ApplicableTimezone   *string                `json:"applicable_timezone,omitempty"`
	BillingConfiguration map[string]interface{} `json:"billing_configuration,omitempty"`
	City                 *string                `json:"city,omitempty"`
	Country              *string                `json:"country,omitempty"`
	CreatedAt            *string                `json:"created_at,omitempty"`
	Currency             *string                `json:"currency,omitempty"`
	Email                *string                `json:"email,omitempty"`
	ExternalID           *string                `json:"external_id,omitempty"`
	LagoID               *string                `json:"lago_id,omitempty"`
	LagoURL              *string                `json:"lago_url,omitempty"`
	LegalName            *string                `json:"legal_name,omitempty"`
	LegalNumber          *string                `json:"legal_number,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Phone                *string                `json:"phone,omitempty"`
	SequentialID         *int64                 `json:"sequential_id,omitempty"`
	Slug                 *string                `json:"slug,omitempty"`
	State                *string                `json:"state,omitempty"`
	Timezone             *string                `json:"timezone,omitempty"`
	URL                  *string                `json:"url,omitempty"`
	Zipcode              *string                `json:"zipcode,omitempty"`
}
