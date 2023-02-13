package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindPlanPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type FindPlanRequest struct {
	PathParams FindPlanPathParams
}

type FindPlanResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Plan                    *shared.Plan
	StatusCode              int64
}
