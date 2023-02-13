package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindCreditNotePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindCreditNoteRequest struct {
	PathParams FindCreditNotePathParams
}

type FindCreditNoteResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	CreditNote              *shared.CreditNote
	StatusCode              int64
}
