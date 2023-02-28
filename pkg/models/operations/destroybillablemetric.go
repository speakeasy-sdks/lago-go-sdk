package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DestroyBillableMetricPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type DestroyBillableMetricRequest struct {
	PathParams DestroyBillableMetricPathParams
}

type DestroyBillableMetricResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	BillableMetric          *shared.BillableMetric
	ContentType             string
	StatusCode              int
}
