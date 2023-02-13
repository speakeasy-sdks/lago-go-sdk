package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindCustomerCurrentUsagePathParams struct {
	CustomerExternalID string `pathParam:"style=simple,explode=false,name=customer_external_id"`
}

type FindCustomerCurrentUsageQueryParams struct {
	ExternalSubscriptionID string `queryParam:"style=form,explode=true,name=external_subscription_id"`
}

type FindCustomerCurrentUsageRequest struct {
	PathParams  FindCustomerCurrentUsagePathParams
	QueryParams FindCustomerCurrentUsageQueryParams
}

type FindCustomerCurrentUsageResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	CustomerUsage           *shared.CustomerUsage
	StatusCode              int64
}
