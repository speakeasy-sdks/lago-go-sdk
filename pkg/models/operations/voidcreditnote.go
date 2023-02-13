package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type VoidCreditNotePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type VoidCreditNoteRequest struct {
	PathParams VoidCreditNotePathParams
}

type VoidCreditNoteResponse struct {
	APIResponseNotAllowed   *shared.APIResponseNotAllowed
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	CreditNote              *shared.CreditNote
	StatusCode              int64
}
