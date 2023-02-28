package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdatePlanPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type UpdatePlanRequest struct {
	PathParams UpdatePlanPathParams
	Request    shared.PlanInput `request:"mediaType=application/json"`
}

type UpdatePlanResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	Plan                           *shared.Plan
	StatusCode                     int
}
