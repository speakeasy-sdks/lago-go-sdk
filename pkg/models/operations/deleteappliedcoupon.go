package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DeleteAppliedCouponPathParams struct {
	AppliedCouponID    string `pathParam:"style=simple,explode=true,name=applied_coupon_id"`
	CustomerExternalID string `pathParam:"style=simple,explode=false,name=customer_external_id"`
}

type DeleteAppliedCouponRequest struct {
	PathParams DeleteAppliedCouponPathParams
}

type DeleteAppliedCouponResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	AppliedCoupon           *shared.AppliedCoupon
	ContentType             string
	StatusCode              int
}
