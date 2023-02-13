package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllCouponsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllCouponsRequest struct {
	QueryParams FindAllCouponsQueryParams
}

type FindAllCouponsResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Coupons                 *shared.Coupons
	StatusCode              int64
}
