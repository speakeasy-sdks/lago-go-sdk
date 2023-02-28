package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateOrganizationRequest struct {
	Request shared.OrganizationInput `request:"mediaType=application/json"`
}

type UpdateOrganizationResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	Organization                   *shared.Organization
	StatusCode                     int
}
