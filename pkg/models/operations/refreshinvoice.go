package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type RefreshInvoicePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type RefreshInvoiceRequest struct {
	PathParams RefreshInvoicePathParams
}

type RefreshInvoiceResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Invoice                 *shared.Invoice
	StatusCode              int64
}
