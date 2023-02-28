package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindCustomerPathParams struct {
	ExternalID string `pathParam:"style=simple,explode=false,name=external_id"`
}

type FindCustomerRequest struct {
	PathParams FindCustomerPathParams
}

type FindCustomerResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Customer                *shared.Customer
	StatusCode              int
}
