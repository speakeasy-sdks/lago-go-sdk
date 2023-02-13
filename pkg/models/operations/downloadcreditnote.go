package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type DownloadCreditNotePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DownloadCreditNoteRequest struct {
	PathParams DownloadCreditNotePathParams
}

type DownloadCreditNoteResponse struct {
	APIResponseNotFound     *shared.APIResponseNotFound
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	CreditNote              *shared.CreditNote
	StatusCode              int64
}
