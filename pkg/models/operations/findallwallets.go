package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllWalletsQueryParams struct {
	ExternalCustomerID string `queryParam:"style=form,explode=true,name=external_customer_id"`
	Page               *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage            *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllWalletsRequest struct {
	QueryParams FindAllWalletsQueryParams
}

type FindAllWalletsResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	Wallets                 *shared.Wallets
}
