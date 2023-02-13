package shared

type OrganizationInputOrganization struct {
	AddressLine1         *string                           `json:"address_line1,omitempty"`
	AddressLine2         *string                           `json:"address_line2,omitempty"`
	BillingConfiguration *BillingConfigurationOrganization `json:"billing_configuration,omitempty"`
	City                 *string                           `json:"city,omitempty"`
	Country              *string                           `json:"country,omitempty"`
	Email                *string                           `json:"email,omitempty"`
	LegalName            *string                           `json:"legal_name,omitempty"`
	LegalNumber          *string                           `json:"legal_number,omitempty"`
	State                *string                           `json:"state,omitempty"`
	Timezone             *string                           `json:"timezone,omitempty"`
	WebhookURL           *string                           `json:"webhook_url,omitempty"`
	Zipode               *string                           `json:"zipode,omitempty"`
}

type OrganizationInput struct {
	Organization *OrganizationInputOrganization `json:"organization,omitempty"`
}
