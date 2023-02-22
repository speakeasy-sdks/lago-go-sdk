package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DeleteCustomerPathParams struct {
	ExternalID string `pathParam:"style=simple,explode=false,name=external_id"`
}

type DeleteCustomerRequest struct {
	PathParams DeleteCustomerPathParams
}

type DeleteCustomerResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Customer                *shared.Customer
	StatusCode              int
}
