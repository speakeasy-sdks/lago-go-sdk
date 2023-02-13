package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateAddOnPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type UpdateAddOnRequest struct {
	PathParams UpdateAddOnPathParams
	Request    shared.AddOnInput `request:"mediaType=application/json"`
}

type UpdateAddOnResponse struct {
	AddOn                          *shared.AddOn
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int64
}
