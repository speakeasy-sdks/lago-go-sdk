package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateAddOnRequest struct {
	Request shared.AddOnInput `request:"mediaType=application/json"`
}

type CreateAddOnResponse struct {
	AddOn                          *shared.AddOn
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int64
}
