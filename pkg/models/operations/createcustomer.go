package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateCustomerRequest struct {
	Request shared.CustomerInput `request:"mediaType=application/json"`
}

type CreateCustomerResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	Customer                       *shared.Customer
	StatusCode                     int64
}
