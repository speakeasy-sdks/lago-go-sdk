package shared

type OrganizationObject struct {
	AddressLine1         *string                           `json:"address_line1,omitempty"`
	AddressLine2         *string                           `json:"address_line2,omitempty"`
	BillingConfiguration *BillingConfigurationOrganization `json:"billing_configuration,omitempty"`
	City                 *string                           `json:"city,omitempty"`
	Country              *string                           `json:"country,omitempty"`
	CreatedAt            *string                           `json:"created_at,omitempty"`
	Email                *string                           `json:"email,omitempty"`
	LagoID               *string                           `json:"lago_id,omitempty"`
	LegalName            *string                           `json:"legal_name,omitempty"`
	LegalNumber          *string                           `json:"legal_number,omitempty"`
	Name                 *string                           `json:"name,omitempty"`
	State                *string                           `json:"state,omitempty"`
	Timezone             *string                           `json:"timezone,omitempty"`
	WebhookURL           *string                           `json:"webhook_url,omitempty"`
	Zipode               *string                           `json:"zipode,omitempty"`
}
