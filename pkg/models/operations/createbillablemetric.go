package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateBillableMetricRequest struct {
	Request shared.BillableMetricInput `request:"mediaType=application/json"`
}

type CreateBillableMetricResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	BillableMetric                 *shared.BillableMetric
	ContentType                    string
	StatusCode                     int
}
