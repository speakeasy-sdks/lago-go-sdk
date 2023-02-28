package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateBillableMetricPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type UpdateBillableMetricRequest struct {
	PathParams UpdateBillableMetricPathParams
	Request    shared.BillableMetricInput `request:"mediaType=application/json"`
}

type UpdateBillableMetricResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	BillableMetric                 *shared.BillableMetric
	ContentType                    string
	StatusCode                     int
}
