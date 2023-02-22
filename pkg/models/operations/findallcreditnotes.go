package operations

import (
	"github.com/speakeasy-sdks/lago-go-sdk/pkg/models/shared"
)

type FindAllCreditNotesQueryParams struct {
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
	Page               *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage            *int64  `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllCreditNotesRequest struct {
	QueryParams FindAllCreditNotesQueryParams
}

type FindAllCreditNotesResponse struct {
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	CreditNotes             *shared.CreditNotes
	StatusCode              int
}
