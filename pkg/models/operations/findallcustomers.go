package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllCustomersQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllCustomersRequest struct {
	QueryParams FindAllCustomersQueryParams
}

type FindAllCustomersResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Customers               *shared.Customers
	StatusCode              int64
}
