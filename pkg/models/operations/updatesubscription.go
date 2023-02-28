package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateSubscriptionPathParams struct {
	ExternalID string `pathParam:"style=simple,explode=false,name=external_id"`
}

type UpdateSubscriptionRequest struct {
	PathParams UpdateSubscriptionPathParams
	Request    shared.SubscriptionUpdateInput `request:"mediaType=application/json"`
}

type UpdateSubscriptionResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int
	Subscription                   *shared.Subscription
}
