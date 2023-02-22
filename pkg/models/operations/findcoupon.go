package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindCouponPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type FindCouponRequest struct {
	PathParams FindCouponPathParams
}

type FindCouponResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Coupon                  *shared.Coupon
	StatusCode              int
}
