package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindInvoicePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindInvoiceRequest struct {
	PathParams FindInvoicePathParams
}

type FindInvoiceResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Invoice                 *shared.Invoice
	StatusCode              int
}
