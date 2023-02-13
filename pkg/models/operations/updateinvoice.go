package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateInvoicePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateInvoiceRequest struct {
	PathParams UpdateInvoicePathParams
	Request    shared.InvoiceInput `request:"mediaType=application/json"`
}

type UpdateInvoiceResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	Invoice                        *shared.Invoice
	StatusCode                     int64
}
