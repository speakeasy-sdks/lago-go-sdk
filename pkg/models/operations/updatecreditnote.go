package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type UpdateCreditNotePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCreditNoteRequest struct {
	PathParams UpdateCreditNotePathParams
	Request    shared.CreditNoteUpdateInput `request:"mediaType=application/json"`
}

type UpdateCreditNoteResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseNotFound            *shared.APIResponseNotFound
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	CreditNote                     *shared.CreditNote
	StatusCode                     int
}
