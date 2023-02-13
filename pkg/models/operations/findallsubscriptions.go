package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllSubscriptionsQueryParams struct {
	ExternalCustomerID string `queryParam:"style=form,explode=true,name=external_customer_id"`
	Page               *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage            *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllSubscriptionsRequest struct {
	QueryParams FindAllSubscriptionsQueryParams
}

type FindAllSubscriptionsResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int64
	Subscriptions           *shared.Subscriptions
}
