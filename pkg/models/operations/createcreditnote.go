package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type CreateCreditNoteRequest struct {
	Request shared.CreditNoteInput `request:"mediaType=application/json"`
}

type CreateCreditNoteResponse struct {
	APIResponseBadRequest          *shared.APIResponseBadRequest
	APIResponseUnauthorized        *shared.APIResponseUnauthorized
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	CreditNote                     *shared.CreditNote
	StatusCode                     int
}
