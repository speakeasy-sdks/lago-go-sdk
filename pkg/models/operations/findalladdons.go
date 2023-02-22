package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllAddOnsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllAddOnsRequest struct {
	QueryParams FindAllAddOnsQueryParams
}

type FindAllAddOnsResponse struct {
	AddOns                  *shared.AddOns
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
}
