package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DestroyPlanPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type DestroyPlanRequest struct {
	PathParams DestroyPlanPathParams
}

type DestroyPlanResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Plan                    *shared.Plan
	StatusCode              int
}
