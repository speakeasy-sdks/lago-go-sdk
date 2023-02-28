package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateCouponPathParams struct {
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type UpdateCouponRequest struct {
	PathParams UpdateCouponPathParams
	Request    shared.CouponInput `request:"mediaType=application/json"`
}

type UpdateCouponResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	Coupon                         *shared.Coupon
	StatusCode                     int
}
