package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllBillableMetricsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllBillableMetricsRequest struct {
	QueryParams FindAllBillableMetricsQueryParams
}

type FindAllBillableMetricsResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	BillableMetrics         *shared.BillableMetrics
	ContentType             string
	StatusCode              int
}
