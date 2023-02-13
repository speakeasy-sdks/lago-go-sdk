package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FinalizeInvoicePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FinalizeInvoiceRequest struct {
	PathParams FinalizeInvoicePathParams
}

type FinalizeInvoiceResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Invoice                 *shared.Invoice
	StatusCode              int64
}
