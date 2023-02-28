package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type ApplyAddOnRequest struct {
	Request shared.AppliedAddOnInput `request:"mediaType=application/json"`
}

type ApplyAddOnResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	AppliedAddOn                   *shared.AppliedAddOn
	ContentType                    string
	StatusCode                     int
}
