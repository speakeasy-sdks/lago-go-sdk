package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllBillableMetricGroupsPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type FindAllBillableMetricGroupsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllBillableMetricGroupsRequest struct {
	PathParams  FindAllBillableMetricGroupsPathParams
	QueryParams FindAllBillableMetricGroupsQueryParams
}

type FindAllBillableMetricGroupsResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Groups                  *shared.Groups
	StatusCode              int64
}
