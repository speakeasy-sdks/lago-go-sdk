package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DownloadInvoicePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DownloadInvoiceRequest struct {
	PathParams DownloadInvoicePathParams
}

type DownloadInvoiceResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	Invoice                 *shared.Invoice
	StatusCode              int
}
