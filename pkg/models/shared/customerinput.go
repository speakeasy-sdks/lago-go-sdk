package shared

type CustomerInputCustomer struct {
	AddressLine1         *string                `json:"address_line1,omitempty"`
	AddressLine2         *string                `json:"address_line2,omitempty"`
	BillingConfiguration map[string]interface{} `json:"billing_configuration,omitempty"`
	City                 *string                `json:"city,omitempty"`
	Country              *string                `json:"country,omitempty"`
	Currency             *string                `json:"currency,omitempty"`
	Email                *string                `json:"email,omitempty"`
	ExternalID           *string                `json:"external_id,omitempty"`
	LagoURL              *string                `json:"lago_url,omitempty"`
	LegalName            *string                `json:"legal_name,omitempty"`
	LegalNumber          *string                `json:"legal_number,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Phone                *string                `json:"phone,omitempty"`
	State                *string                `json:"state,omitempty"`
	Timezone             *string                `json:"timezone,omitempty"`
	URL                  *string                `json:"url,omitempty"`
	Zipode               *string                `json:"zipode,omitempty"`
}

type CustomerInput struct {
	Customer *CustomerInputCustomer `json:"customer,omitempty"`
}
