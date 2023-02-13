package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateSubscriptionRequest struct {
	Request shared.SubscriptionCreateInput `request:"mediaType=application/json"`
}

type CreateSubscriptionResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int64
	Subscription                   *shared.Subscription
}
