package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DestroyCouponPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type DestroyCouponRequest struct {
	PathParams DestroyCouponPathParams
}

type DestroyCouponResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Coupon                  *shared.Coupon
	StatusCode              int
}
