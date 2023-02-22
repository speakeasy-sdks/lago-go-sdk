package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type ApplyCouponRequest struct {
	Request shared.AppliedCouponInput `request:"mediaType=application/json"`
}

type ApplyCouponResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	AppliedCoupon                  *shared.AppliedCoupon
	ContentType                    string
	StatusCode                     int
}
