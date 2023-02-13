package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindBillableMetricPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type FindBillableMetricRequest struct {
	PathParams FindBillableMetricPathParams
}

type FindBillableMetricResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	BillableMetric          *shared.BillableMetric
	ContentType             string
	StatusCode              int64
}
