package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DestroySubscriptionPathParams struct {
	ExternalID string `pathParam:"style=simple,explode=false,name=external_id"`
}

type DestroySubscriptionRequest struct {
	PathParams DestroySubscriptionPathParams
}

type DestroySubscriptionResponse struct {
	APIResponseNotAllowed   *shared.APIResponseNotAllowed
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	Subscription            *shared.Subscription
}
