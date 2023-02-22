package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateCouponRequest struct {
	Request shared.CouponInput `request:"mediaType=application/json"`
}

type CreateCouponResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	Coupon                         *shared.Coupon
	StatusCode                     int
}
