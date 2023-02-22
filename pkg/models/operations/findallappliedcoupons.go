package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllAppliedCouponsStatusEnum string

const (
	FindAllAppliedCouponsStatusEnumActive     FindAllAppliedCouponsStatusEnum = "active"
	FindAllAppliedCouponsStatusEnumTerminated FindAllAppliedCouponsStatusEnum = "terminated"
)

type FindAllAppliedCouponsQueryParams struct {
	ExternalCustomerID *string                          `queryParam:"style=form,explode=true,name=external_customer_id"`
	Page               *int64                           `queryParam:"style=form,explode=true,name=page"`
	PerPage            *int64                           `queryParam:"style=form,explode=true,name=per_page"`
	Status             *FindAllAppliedCouponsStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type FindAllAppliedCouponsRequest struct {
	QueryParams FindAllAppliedCouponsQueryParams
}

type FindAllAppliedCouponsResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	AppliedCoupons          *shared.AppliedCoupons
	ContentType             string
	StatusCode              int
}
